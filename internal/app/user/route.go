package user

import (
	"compass_mini_api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/change_password/:id", h.ChangePassword, middleware.Authentication)
}
