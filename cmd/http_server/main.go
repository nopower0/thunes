package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"thunes/bindings"
	"thunes/components"
	"thunes/handlers/http"
	"thunes/middlewares"
	"thunes/objects/models"
	"thunes/settings"
	"thunes/tools"
)

func errorHandler(e error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	logger := tools.GetLogger(c)
	logger.Error("api", zap.Error(e))
	_ = bindings.JSONResponse(c, bindings.ServerError)
}

func main() {
	settings.Init()
	models.Init()
	components.Init()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = errorHandler
	e.Use(
		middlewares.PrometheusMiddleware(),
		middlewares.LoggerMiddleware(),
		middlewares.BodyDumpMiddleware(),
	)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	http.SetupRouter(e)

	if err := e.Start(fmt.Sprintf(":%d", settings.Port)); err != nil {
		zap.L().Fatal("error starting Echo server", zap.Int("port", settings.Port), zap.Error(err))
	}
}
