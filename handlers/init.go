package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/external"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/config"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/middleware"

	_ "github.com/p-jirayusakul/go-flat-arch-template/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerHttpHandler struct {
	cfg   *config.Config
	store database.Store
	exApi external.ExternalAPI
}

func NewHandler(
	app *echo.Echo,
	cfg *config.Config,
	store database.Store,
	exApi external.ExternalAPI,
) *ServerHttpHandler {

	handler := &ServerHttpHandler{
		cfg:   cfg,
		store: store,
		exApi: exApi,
	}

	// auth
	app.GET(common.DOCS_URL+"/*", echoSwagger.WrapHandler)

	authGroup := app.Group(common.BASE_URL + "/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)

	// address
	addressesGroup := app.Group(common.BASE_URL + "/profile")
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
		return fmt.Errorf(common.ErrAccountIsInvalid.Error())
	}

	return nil
}
