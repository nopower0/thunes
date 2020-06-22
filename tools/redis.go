package tools

import (
	"github.com/go-redis/redis"
)

type RedisClientConf struct {
	RW string `json:"rw"`
	RO string `json:"ro"`
}

type RedisClient struct {
	RW *redis.Client
	RO *redis.Client
}

func NewRedisClient(conf *RedisClientConf) (*RedisClient, error) {
	client := new(RedisClient)
	if options, err := redis.ParseURL(conf.RW); err != nil {
		return nil, err
	} else {
		client.RW = redis.NewClient(options)
	}
	if len(conf.RO) == 0 {
		client.RO = client.RW
	} else {
		if options, err := redis.ParseURL(conf.RO); err != nil {
			return nil, err
		} else {
			client.RO = redis.NewClient(options)
		}
	}
	return client, nil
}
