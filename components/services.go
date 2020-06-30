package components

import (
	"go.uber.org/zap"
	"thunes/settings"
	"thunes/tools"
)

var (
	DefaultAuthService IAuthService
)

func Init() {
	r, err := tools.NewRedisClient(settings.DefaultRedisConf)
	if err != nil {
		zap.L().Fatal("Invalid Redis Conf", zap.Any("conf", settings.DefaultRedisConf), zap.Error(err))
	}
	DefaultAuthService = NewAuthService(r)
}
