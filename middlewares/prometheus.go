package middlewares

import (
	prometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
)

func PrometheusMiddleware() echo.MiddlewareFunc {
	config := prometheus.DefaultConfig
	config.Namespace = "thunes"
	config.Subsystem = ""
	return prometheus.MetricsMiddlewareWithConfig(config)
}
