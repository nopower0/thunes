package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"thunes/tools"
)

func bodyDumpHandler(c echo.Context, reqBody []byte, rspBody []byte) {
	logger := tools.GetLogger(c)
	logger.Info("api", zap.Any("data", map[string]string{
		"req_data": string(reqBody),
		"rsp_data": string(rspBody),
	}))
}

func BodyDumpMiddleware() echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(_ echo.Context) bool {
			return false
		},
		Handler: bodyDumpHandler,
	})
}
