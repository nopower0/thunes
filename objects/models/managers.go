package models

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"thunes/settings"
	"time"
	"xorm.io/xorm"
)

var (
	DefaultWalletManager *WalletManager
	DefaultUserManager   *UserManager
)

func Init() {
	engine, err := xorm.NewEngine("mysql", settings.DefaultDB)
	if err != nil {
		zap.L().Fatal("error creating default engine", zap.String("db", settings.DefaultDB), zap.Error(err))
	} else {
		engine.SetConnMaxLifetime(1 * time.Minute)
		engine.SetMaxOpenConns(100)
		engine.SetMaxIdleConns(25)
	}

	DefaultWalletManager = NewWalletManager(engine)
	DefaultUserManager = NewUserManager(engine)
}
