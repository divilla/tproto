package auth_controller

import (
	"github.com/divilla/tproto/users/internal/auth/auth_service"
	"github.com/divilla/tproto/users/internal/containers"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service *auth_service.Service
}

func NewV1(g *echo.Group, mc *containers.Main) {
	c := &controller{
		service: auth_service.New(mc),
	}

	g = g.Group("/v1")
	g = g.Group("/auth")
	g.POST("/login", c.login)
}

func (c *controller) login(ctx echo.Context) error {
	var u auth_service.LoginDTO
	if err := ctx.Bind(&u); err != nil {
		return err
	}

	if user, err := c.service.Login(&u); err != nil {
		return err
	} else {
		return ctx.JSON(200, user)
	}
}
