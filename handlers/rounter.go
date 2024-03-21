package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/middleware"
	"github.com/p-jirayusakul/go-flat-arch-template/utils"
)

type ServerHttpHandler struct {
	cfg   *utils.Config
	store database.Store
}

func NewHandler(
	app *echo.Echo,
	cfg *utils.Config,
	store database.Store,
) *ServerHttpHandler {

	handler := &ServerHttpHandler{
		cfg:   cfg,
		store: store,
	}

	// auth
	var baseAPI = "/api/v1"
	authGroup := app.Group(baseAPI + "/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)

	// address
	addressesGroup := app.Group(baseAPI + "/profile")
	addressesGroup.Use(middleware.ConfigJWT(cfg.JWT_SECRET))
	addressesGroup.POST("/addresses", handler.CreateAddresses)
	addressesGroup.GET("/addresses", handler.ListAddresses)
	addressesGroup.PUT("/addresses/:id", handler.UpdateAddresses)
	addressesGroup.DELETE("/addresses/:id", handler.DeleteAddresses)

	return handler
}

// utils function
func (s *ServerHttpHandler) GetTokenID(c echo.Context) error {
	isAlreadyExists, err := s.store.IsAccountAlreadyExists(c.Request().Context(), c.Get("accountsID").(string))
	if err != nil {
		return err
	}

	if !isAlreadyExists {
		return fmt.Errorf(utils.ErrAccountIsInvalid.Error())
	}

	return nil
}
