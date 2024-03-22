package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/request"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/response"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/middleware"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/utils"
)

// Register
// @Summary      Register By email and password
// @Description  register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body request.RegisterRequest true "body request"
// @Success      201  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/auth/register [post]
func (s *ServerHttpHandler) Register(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.RegisterRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	isEmailAlready, err := s.store.IsEmailAlreadyExists(ctx, body.Email)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	if isEmailAlready {
		return utils.RespondWithError(http.StatusBadRequest, common.ErrEmailIsAlreadyExists.Error())
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	params := database.CreateAccountParams{
		Email:    body.Email,
		Password: hashedPassword,
	}

	// Save to Repository
	_, err = s.store.CreateAccount(ctx, params)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Test Call APIs External Project
	resultAPIs, err := s.exApi.GetPosts()
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("resultAPIs", resultAPIs)

	// Response
	var payload interface{}
	message := "registration completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}

// Login
// @Summary      Login By email and password
// @Description  register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body request.LoginRequest true "body request"
// @Success      200  {object}  utils.SuccessResponse.Data{data=response.LoginResponse}
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/auth/login [post]
func (s *ServerHttpHandler) Login(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.LoginRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	account, err := s.store.GetAccountByEmail(ctx, body.Email)
	if err != nil {
		if errors.Is(err, common.ErrDBNoRows) {
			return utils.RespondWithError(http.StatusUnauthorized, common.ErrLoginFail.Error())
		}
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	err = utils.CheckPassword(body.Password, account.Password)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, common.ErrLoginFail.Error())
	}

	token, err := middleware.CreateToken(middleware.CreateTokenDTO{
		UserID:    account.ID,
		Secret:    s.cfg.JWT_SECRET,
		ExpiresAt: time.Now().Add(time.Hour * 72),
	})
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	payload := response.LoginResponse{
		Token: token,
	}
	message := "login completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}
