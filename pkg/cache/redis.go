package cache

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/kyh0703/go-project-layout/configs"
)

var client *redis.Client

func ProvideRedisClient() (*redis.Client, error) {
	// create new client
	client = redis.NewClient(&redis.Options{Addr: configs.Env.CacheUrl})
	// ping test
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}

var RedisSet = wire.NewSet(ProvideRedisClient)
