package redis

import (
	"urls/config"
	"github.com/go-redis/redis/v8"
)

func InitRedisConnect() redis.Client {
	cfg, _ := config.LoadConfig("../config")
	rdb := redis.NewClient(&redis.Options{
		Addr:	  cfg.Redis.Address,
		Password: cfg.Redis.Password, // no password set
		DB:		  cfg.Redis.DB,  // use default DB
	})
	return *rdb
}