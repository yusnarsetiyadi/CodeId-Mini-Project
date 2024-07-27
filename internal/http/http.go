package http

import (
	"fmt"
	"net/http"

	_ "compass_mini_api/docs"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	"compass_mini_api/internal/app/auth"
	"compass_mini_api/internal/app/company"
	"compass_mini_api/internal/app/employee"
	"compass_mini_api/internal/app/feature"
	"compass_mini_api/internal/app/parameteritem"
	"compass_mini_api/internal/app/user"
	"compass_mini_api/internal/config"

	"compass_mini_api/internal/factory"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = config.Get().Server.App
		VERSION = config.Get().Server.Version
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// logs
	e.Static("/logs", "/pkg/log/output")
	// routes
	api_v1 := e.Group("/api/v1")
	auth.NewHandler(f).Route(api_v1.Group("/auth"))
	feature.NewHandler(f).Route(api_v1.Group("/feature"))
	user.NewHandler(f).Route(api_v1.Group("/user"))
	parameteritem.NewHandler(f).Route(api_v1.Group("/parameteritem"))
	company.NewHandler(f).Route(api_v1.Group("/company"))
	employee.NewHandler(f).Route(api_v1.Group("/employee"))

}
