package bootstrap

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/MikelSot/amaris-beer/model"
)

func newRedisClient(ctx context.Context, config model.RedisConfig) *redis.Client {
	db, err := strconv.Atoi(config.Name)
	if err != nil {
		log.Fatalf("No se pudo convertir el valor de REDIS_DB a int: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       db,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("No se pudo conectar a Redis: %v", err)
	}

	return rdb
}
