package containers

import (
	"github.com/divilla/tproto/users/pkg/random_bytes"
	"github.com/labstack/echo/v4"
)

type Main struct {
	log echo.Logger
	rbs *random_bytes.Service
}

func NewMain(l echo.Logger) *Main {
	return &Main{
		log: l,
		rbs: random_bytes.New(l),
	}
}

func (m *Main) Log() echo.Logger {
	return m.log
}

func (m *Main) RandomBytesService() *random_bytes.Service {
	return m.rbs
}
