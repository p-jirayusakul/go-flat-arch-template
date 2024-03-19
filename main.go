package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers"
	"github.com/p-jirayusakul/go-flat-arch-template/middleware"
	"github.com/p-jirayusakul/go-flat-arch-template/utils"
)

const PATH_CONFIG = ".env"

var (
	config = utils.InitConfigs(PATH_CONFIG)
	db     = database.InitDatabase(config)
)

func main() {

	// plug database
	store := database.NewStore(db)

	// plug controller
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Validator = middleware.NewCustomValidator()
	app.Use(middleware.ErrorHandler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(middleware.LogHandler(logger))

	handlers.NewHandler(app, &config, store)
	app.Logger.Fatal(app.Start(":" + config.API_PORT))
}
