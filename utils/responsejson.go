package utils

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"something went wrong"`
}

func RespondWithError(c echo.Context, code int, message string) error {
	return c.JSON(code, ErrorResponse{Message: message, Status: "error"})
}

func RespondWithJSON(c echo.Context, code int, message string, payload interface{}) error {
	return c.JSON(code, SuccessResponse{Message: message, Status: "success", Data: payload})
}
