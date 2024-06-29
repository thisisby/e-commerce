package caches

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type RedisCache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
	Incr(key string) error
}

type redisCache struct {
	host     string
	db       int
	password string
	expires  time.Duration
	client   *redis.Client
}

func NewRedisCache(host string, db int, password string, expires time.Duration) RedisCache {
	return &redisCache{
		host:     host,
		db:       db,
		password: password,
		expires:  expires,
		client: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		}),
	}
}

func (c *redisCache) Set(key string, value any) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, json, c.expires*time.Minute).Err()
}

func (c *redisCache) Get(key string) (any, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var result any
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *redisCache) Delete(key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *redisCache) Incr(key string) error {
	return c.client.Incr(ctx, key).Err()
}
