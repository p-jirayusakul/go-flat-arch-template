package handlers

import (
	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	config "github.com/p-jirayusakul/go-flat-arch-template/utils"
)

type ServerHttpHandler struct {
	cfg   *config.Config
	store database.Store
}

func NewHandler(
	app *echo.Echo,
	cfg *config.Config,
	store database.Store,
) *ServerHttpHandler {

	handler := &ServerHttpHandler{
		cfg:   cfg,
		store: store,
	}

	// auth
	authGroup := app.Group("/api/v1/auth")
	authGroup.POST("/register", handler.Register)

	return handler
}
