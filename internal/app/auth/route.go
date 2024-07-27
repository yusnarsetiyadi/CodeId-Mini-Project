package auth

import (
	"compass_mini_api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/login", h.Login)
	g.POST("/splash", h.Splash)
	g.DELETE("/logout", h.Logout, middleware.Authentication)
	g.GET("/get_data_token", h.GetDataToken, middleware.Authentication)
}
