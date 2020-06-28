package settings

import (
	"go.uber.org/zap"
	"os"
	"thunes/tools"
	"time"
)

var (
	Port int

	DefaultDB        string
	DefaultRedisConf *tools.RedisClientConf
	TokenTTL         time.Duration
)

type config struct {
	Port int `json:"port"`
	DB   struct {
		Default string `json:"default"`
	} `json:"db"`
	Redis struct {
		Default *tools.RedisClientConf `json:"default"`
	} `json:"redis"`
	Token struct {
		TTL int `json:"ttl"`
	} `json:"token"`
}

func Init() {
	initLogger()

	var confPath = "conf/dev.json"
	if path, exist := os.LookupEnv("CONF"); exist {
		confPath = path
	}

	conf := new(config)
	if err := tools.LoadConfigFromFile(confPath, conf); err != nil {
		zap.L().Fatal("error loading config file", zap.String("path", confPath), zap.Error(err))
	}

	Port = conf.Port

	DefaultDB = conf.DB.Default
	DefaultRedisConf = conf.Redis.Default
	TokenTTL = time.Duration(conf.Token.TTL) * time.Second
}
