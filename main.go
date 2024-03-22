package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/external"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/config"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/middleware"
)

const PATH_CONFIG = ".env"

var (
	cfg = config.InitConfigs(PATH_CONFIG)
	db  = database.InitDatabase(cfg)
)

// @title           Flat Architecture
// @version         1.0
// @description     This is template flat arch

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @host      localhost:3000
func main() {

	// plug external APIs
	exApi := external.New(&cfg)

	// plug database
	store := database.NewStore(db)

	// plug controller
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// init validator
	app.Validator = middleware.NewCustomValidator()
	app.Use(middleware.ErrorHandler)

	// init log
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(middleware.LogHandler(logger))

	// add all plug to handler
	handlers.NewHandler(app, &cfg, store, exApi)
	app.Logger.Fatal(app.Start(":" + cfg.API_PORT))
}
