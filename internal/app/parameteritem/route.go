package parameteritem

import (
	"compass_mini_api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("/get_all_division", h.GetAllDivision, middleware.Authentication)
}
