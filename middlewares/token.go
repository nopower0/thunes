package middlewares

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"thunes/bindings"
	"thunes/components"
	"thunes/settings"
	"thunes/tools"
)

func TokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(settings.HeaderToken)
			if len(token) == 0 {
				return bindings.JSONResponse(c, bindings.TokenError)
			}
			if info, err := components.DefaultAuthService.GetTokenInfo(token); err != nil {
				return err
			} else if info == nil {
				return bindings.JSONResponse(c, bindings.TokenError)
			} else {
				tools.AttachTokenInfo(c, info)
				tools.AttachLogger(c, tools.GetLogger(c).With(zap.Any("token", info)))
			}
			return next(c)
		}
	}
}
