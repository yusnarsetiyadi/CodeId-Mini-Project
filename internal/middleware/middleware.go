package middleware

import (
	"fmt"
	"net/http"
	"os"

	"compass_mini_api/internal/factory"
	"compass_mini_api/pkg/util/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
	"github.com/sirupsen/logrus"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP  = os.Getenv("APP")
		ENV  = os.Getenv("ENV")
		NAME = fmt.Sprintf("%s-%s", APP, ENV)
	)

	e.Use(Context, Trace)
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "x-user-id", "ngrok-skip-browser-warning", echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodHead, http.MethodPatch},
		}),
		echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
			LogLatency:   true,
			LogRemoteIP:  true,
			LogHost:      true,
			LogMethod:    true,
			LogURI:       true,
			LogUserAgent: true,
			LogStatus:    true,
			LogError:     true,
			LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
				colorer := color.New()
				statusColor := colorer.Green(v.Status)
				typeLog := "INFO"
				switch {
				case v.Status >= 500:
					statusColor = colorer.Red(v.Status)
					typeLog = "ERR"
				case v.Status >= 400:
					statusColor = colorer.Yellow(v.Status)
					typeLog = "WARN"
				case v.Status >= 300:
					statusColor = colorer.Cyan(v.Status)
					typeLog = "INFO"
				}
				data := fmt.Sprintf("\n| %s | Host: %s | Status: %s | LatencyHuman: %d | UserAgent: %s | RemoteIp: %s | Method: %s | Uri: %s |\n", NAME, v.Host, statusColor, v.Latency, v.UserAgent, v.RemoteIP, v.Method, v.URI)
				if typeLog == "INFO" || v.Status == 200 {
					logrus.Info(data)
				} else if typeLog == "WARN" {
					logrus.Warn(data)
				} else {
					logrus.Error(data)
				}
				return nil
			},
		}),
	)
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}

	newUserAuthService(f)

}
