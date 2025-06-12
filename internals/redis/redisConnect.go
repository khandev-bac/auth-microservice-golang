package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	rds := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisClient{
		Client: rds,
	}
}

func (r *RedisClient) BlackListToken(token string, ttl time.Duration) error {
	return r.Client.Set(ctx, "BlackListToken:"+token, "true", ttl).Err()
}
func (r *RedisClient) IsBlackListToken(token string) (bool, error) {
	res, err := r.Client.Get(ctx, "BlackListToken:"+token).Result()
	if err == redis.Nil {
		return false, nil
	}
	return res == "true", nil
}
