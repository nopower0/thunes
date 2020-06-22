package http

import (
	"github.com/labstack/echo/v4"
	"thunes/middlewares"
)

func SetupRouter(e *echo.Echo) {
	// Token Related
	{
		g := e.Group("/token")
		h := new(TokenHandler)
		g.POST("/request", h.Request)
	}
	// UserRelated
	{
		g := e.Group("/user", middlewares.TokenMiddleware())
		h := new(UserHandler)
		g.POST("/login", h.Login)
	}
	// Wallet Related
	{
		g := e.Group("/wallet", middlewares.TokenMiddleware(), middlewares.AuthMiddleware())
		h := new(WalletHandler)
		g.POST("/get", h.Get)
	}
}
