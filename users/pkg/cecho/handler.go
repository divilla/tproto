package cecho

import (
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(ctx *ccontext) error

func Handler(hf HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*ccontext)
		return hf(ctx)
	}
}
