package cecho

import (
	"github.com/divilla/tproto/users/internal/auth/auth_domain"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewHTTPErrorHandler(e *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {

		if c.Response().Committed {
			return
		}

		he, ok := err.(*echo.HTTPError)
		if ok {
			if he.Internal != nil {
				if herr, ok := he.Internal.(*echo.HTTPError); ok {
					he = herr
				}
			}
		} else {
			he = &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}

		var code int
		var message interface{}
		switch terr := err.(type) {
		case validation.Errors:
			code = http.StatusUnprocessableEntity
			message = terr
		case auth_domain.Errors:
			code = http.StatusUnprocessableEntity
			message = terr
		case *auth_domain.HttpError:
			code = terr.Code
			message = echo.Map{"message": terr.Message}
		default:
			code = he.Code
			message = he.Message
			if m, ok := he.Message.(string); ok {
				if e.Debug {
					message = echo.Map{"message": m, "error": err.Error()}
				} else {
					message = echo.Map{"message": m}
				}
			}
		}

		// Send response
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			e.Logger.Error(err)
		}
	}
}
