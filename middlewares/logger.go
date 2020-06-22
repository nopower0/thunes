package middlewares

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"thunes/tools"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			attachLogger(c)
			return next(c)
		}
	}
}

func attachLogger(c echo.Context) {
	header := c.Request().Header
	logger := zap.L().With(
		zap.Any("metadata", map[string]string{
			"ua":          header.Get("USER-AGENT"),
			"remote_addr": c.RealIP(),
			"uri":         c.Path(),
		}),
	)
	tools.AttachLogger(c, logger)
}
