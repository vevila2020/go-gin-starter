package redis

import (
	"context"
	"fmt"
	"time"

	"go-gin-starter/internal/config"
	"go-gin-starter/internal/logger"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var Client *redis.Client

func Init() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("config not initialized")
	}

	redisCfg := cfg.Redis

	if !redisCfg.Enabled {
		logger.Info("redis is disabled, skipping initialization")
		return nil
	}

	Client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port),
		Password:     redisCfg.Password,
		DB:           redisCfg.DB,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}

	logger.Info("redis connected successfully",
		zap.String("host", redisCfg.Host),
		zap.Int("port", redisCfg.Port),
		zap.Int("db", redisCfg.DB),
	)

	return nil
}

func GetClient() *redis.Client {
	return Client
}

func IsEnabled() bool {
	return Client != nil
}

func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if !IsEnabled() {
		return fmt.Errorf("redis is not enabled")
	}
	return Client.Set(ctx, key, value, expiration).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	if !IsEnabled() {
		return "", fmt.Errorf("redis is not enabled")
	}
	return Client.Get(ctx, key).Result()
}

func Del(ctx context.Context, keys ...string) error {
	if !IsEnabled() {
		return fmt.Errorf("redis is not enabled")
	}
	return Client.Del(ctx, keys...).Err()
}

func Exists(ctx context.Context, keys ...string) (int64, error) {
	if !IsEnabled() {
		return 0, fmt.Errorf("redis is not enabled")
	}
	return Client.Exists(ctx, keys...).Result()
}

func Expire(ctx context.Context, key string, expiration time.Duration) error {
	if !IsEnabled() {
		return fmt.Errorf("redis is not enabled")
	}
	return Client.Expire(ctx, key, expiration).Err()
}
