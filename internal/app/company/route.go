package company

import (
	"compass_mini_api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAllCompany, middleware.Authentication)
}
