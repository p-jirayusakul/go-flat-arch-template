package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/go-flat-arch-template/utils"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// Handle errors here
			switch e := err.(type) {
			case *echo.HTTPError:
				return utils.RespondWithError(c, e.Code, e.Message.(string))
			default:
				return utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}
		return nil
	}
}
