package feature

import (
	"compass_mini_api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("/list", h.GetFeatureList, middleware.Authentication)
	g.GET("/sub/:id", h.GetFeatureSub, middleware.Authentication)
}
