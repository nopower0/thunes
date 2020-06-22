package tools

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"thunes/objects/business"
)

const (
	TokenInfo = "token"
	Logger    = "logger"
)

func AttachTokenInfo(c echo.Context, info *business.TokenInfo) {
	if info == nil {
		return
	}
	c.Set(TokenInfo, info)
}

func GetTokenInfo(c echo.Context) *business.TokenInfo {
	info := c.Get(TokenInfo)
	if info == nil {
		return nil
	}
	return info.(*business.TokenInfo)
}

func AttachLogger(c echo.Context, logger *zap.Logger) {
	c.Set(Logger, logger)
}

func GetLogger(c echo.Context) *zap.Logger {
	logger := c.Get(Logger)
	if logger == nil {
		return nil
	}
	return logger.(*zap.Logger)
}
