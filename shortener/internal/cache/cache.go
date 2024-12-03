package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"shortener/internal/config"
	"strconv"
)

func CacheClient(config config.Cache, log *zap.Logger) *redis.Client {
	log.Info("connecting with the cache")
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Error("failed to connect the cache", zap.Error(err))
		panic(err)
	}

	return client
}
