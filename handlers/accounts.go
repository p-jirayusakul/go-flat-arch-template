package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/request"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/response"
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
	result, err := s.store.CreateAccount(ctx, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	payload := response.RegisterResponse{
		ID: result,
	}
	message := "registration completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}
