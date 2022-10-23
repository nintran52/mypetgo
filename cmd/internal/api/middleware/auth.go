package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nintran52/mypetgo/cmd/internal/api"
)

type AuthMode int

const (
	AuthModeRequired AuthMode = iota
	AuthModeSecure
	AuthModeOptional
	AuthModeTry
	AuthModeNone
)

type AuthConfig struct {
	S       *api.Server
	Mode    AuthMode
	Skipper middleware.Skipper
}

var (
	DefaulAuthConfig = AuthConfig{
		Mode:    AuthModeRequired,
		Skipper: middleware.DefaultSkipper,
	}
)

func AuthWithConfig(config AuthConfig) echo.MiddlewareFunc {
	if config.S == nil {
		panic("auth middleware: server is required")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
