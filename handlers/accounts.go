package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/request"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/response"
	"github.com/p-jirayusakul/go-flat-arch-template/middleware"
	"github.com/p-jirayusakul/go-flat-arch-template/utils"
)

func (s *ServerHttpHandler) Register(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.RegisterRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Logic
	isEmailAlready, err := s.store.IsEmailAlreadyExists(ctx, body.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if isEmailAlready {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrEmailIsAlreadyExists.Error())
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	params := database.CreateAccountParams{
		Email:    body.Email,
		Password: hashedPassword,
	}

	// Save to Repository
	_, err = s.store.CreateAccount(ctx, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "registration completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}

func (s *ServerHttpHandler) Login(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.LoginRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Logic
	account, err := s.store.GetAccountByEmail(ctx, body.Email)
	if err != nil {
		if errors.Is(err, utils.ErrDBNoRows) {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrLoginFail.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = utils.CheckPassword(body.Password, account.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrLoginFail.Error())
	}

	token, err := middleware.CreateToken(middleware.CreateTokenDTO{
		UserID:    account.ID,
		Secret:    s.cfg.JWT_SECRET,
		ExpiresAt: time.Now().Add(time.Hour * 72),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	payload := response.LoginResponse{
		Token: token,
	}
	message := "login completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}
