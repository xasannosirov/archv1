package cache

import (
	"archv1/internal/pkg/config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Cache *redis.Client
}

func NewRedis(cfg *config.Config) (*Redis, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		DB:       cfg.RedisDB,
		Password: cfg.RedisPWD,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil && pong != "PONG" {
		return nil, errors.New("redis ping error")
	}

	return &Redis{
		Cache: client,
	}, nil
}

// Set save key and value to cache for 3 minutes
func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cmd := r.Cache.Set(ctx, key, str, expiration)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	str, err := r.Cache.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return str, nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	cmd := r.Cache.Del(ctx, key)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
