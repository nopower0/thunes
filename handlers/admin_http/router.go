package admin_http

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
	// User Related
	{
		g := e.Group("/user", middlewares.TokenMiddleware())
		h := new(UserHandler)
		g.POST("/login", h.Login)
		g.POST("/add", h.Add)
	}
	// Wallet Related
	{
		g := e.Group("/wallet", middlewares.TokenMiddleware())
		h := new(WalletHandler)
		g.POST("/get_summary", h.GetSummary)
		g.POST("/get_transaction_summary", h.GetTransactionSummary)
		g.POST("/list", h.List)
	}
}
