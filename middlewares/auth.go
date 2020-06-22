package middlewares

import (
	"github.com/labstack/echo/v4"
	"thunes/bindings"
	"thunes/tools"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		 return func(c echo.Context) error {
		 	tokenInfo := tools.GetTokenInfo(c)
		 	if tokenInfo.UID == 0 {
		 		return bindings.JSONResponse(c, bindings.UserNotLoginError)
			}
			return next(c)
		 }
	}
}
