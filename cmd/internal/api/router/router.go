package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nintran52/mypetgo/cmd/internal/api"
	"github.com/nintran52/mypetgo/cmd/internal/api/middleware"
)

func Init(s *api.Server) {
	s.Echo = echo.New()
	s.Router = &api.Router{
		Routes: nil,
		Root:   s.Echo.Group(""),
		APIV1Auth: s.Echo.Group("/api/v1/auth", middleware.AuthWithConfig(middleware.AuthConfig{
			S:    s,
			Mode: middleware.AuthModeRequired,
			Skipper: func(c echo.Context) bool {
				switch c.Path() {
				case "/api/v1/auth/forgot-password",
					"api/v1/auth/login",
					"api/v1/auth/refresh",
					"api/v1/auth/register":
					return true
				}
				return false
			},
		})),
	}
}
