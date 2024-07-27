package employee

import (
	"compass_mini_api/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAllEmployee, middleware.Authentication)
	g.POST("", h.Create, middleware.Authentication)
	g.GET("/:id", h.GetEmployeeById, middleware.Authentication)
	g.PUT("/:id", h.Update, middleware.Authentication)
	g.GET("/supervisor", h.GetAllEmployeeSupervisor, middleware.Authentication)
	g.POST("/with_base64", h.CreateWithBase64, middleware.Authentication)
	g.PUT("/with_base64/:id", h.UpdateWithBase64, middleware.Authentication)
	g.GET("/employeephoto/:base64", h.GetEmployeePhoto, middleware.Authentication)
}
