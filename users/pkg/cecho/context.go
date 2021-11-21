package cecho

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"strconv"
)

var (
	ErrEmptyBody = errors.New("request body does not contain json")
	ErrReadBody = errors.New("failed reading request body")
	ErrCloseBody = errors.New("failed closing request body")
)

type (
	ccontext struct {
		echo.Context
		//identity     *entity.Identity
	}
)

func NewContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &ccontext{
				Context: c,
			}

			return next(ctx)
		}
	}
}

func (c *ccontext) RequestContext() context.Context {
	return c.Request().Context()
}

func (c *ccontext) BodyBytes() ([]byte, error) {
	body, err := c.Request().GetBody()
	if err != nil {
		return nil, ErrEmptyBody
	}

	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, ErrReadBody
	}

	if err = body.Close(); err != nil {
		return nil, ErrCloseBody
	}

	return bodyBytes, nil
}

func (c *ccontext) BodyString() (string, error) {
	if bb, err := c.BodyBytes(); err != nil {
		return "", err
	} else {
		return string(bb), nil
	}
}

func (c *ccontext) BodyMap() (echo.Map, error) {
	bb, err := c.BodyBytes()
	if err != nil {
		return nil, err
	}

	var m echo.Map
	err = jsoniter.Unmarshal(bb, m)
	return m, nil
}

func (c *ccontext) JSONBytes(code int, json []byte) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	_, err := c.Response().Write(json)
	if err != nil {
		return err
	}

	c.Response().Status = code
	return nil
}

func (c *ccontext) JSONString(code int, json string) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.String(code, json)
}

func (c *ccontext) ParamInt64(name string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(c.Param(name), 10, 64)
	if err != nil || value < 1 {
		value = defaultValue
	}

	return value
}

func (c *ccontext) QueryParamInt64(name string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(c.QueryParam(name), 10, 64)
	if err != nil || value < 1 {
		value = defaultValue
	}

	return value
}
