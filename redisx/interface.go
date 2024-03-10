package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v7"
)

// RedisAPI is the public Redis API interface
type RedisAPI interface {
	Connect() error
	// Set set a key/value with an expiry duration, after this duration the key/value is deleted from redis
	Set(ctx context.Context, key string, expiry time.Duration, data interface{}) error
	// HSet ...
	HSet(ctx context.Context, key string, field string, data interface{}) error
	// HDel ...
	HDel(ctx context.Context, key string, fields ...string) error
	// Get ...
	Get(ctx context.Context, key string) (string, error)
	// HGet ...
	HGet(ctx context.Context, key string, field string, response interface{}) error
	// HGetAll ...
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	// Keys return all keys from redis if the pattern is set to "*", not safe, use Scan instead
	Keys(ctx context.Context, pattern string) ([]string, error)
	// Scan returns every keys from a redis database safely
	Scan(ctx context.Context, match string, count int) (map[string]string, error)
	// Incr ...
	Incr(ctx context.Context, key string) (int64, error)
	// RPush ...
	RPush(context.Context, string, ...interface{}) error
	// BLPop ...
	BLPop(context.Context, time.Duration, ...string) ([]string, error)
	// LPop ...
	LPop(context.Context, string) (string, error)
}

// ClientAPI is the internal Client API interface
type ClientAPI interface {
	Ping() *redis.StatusCmd
	Close() error
	HDel(string, ...string) *redis.IntCmd
	Get(string) *redis.StringCmd
	HGet(string, string) *redis.StringCmd
	HGetAll(string) *redis.StringStringMapCmd
	Keys(string) *redis.StringSliceCmd
	Scan(uint64, string, int64) *redis.ScanCmd
	Incr(string) *redis.IntCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	HSet(string, ...interface{}) *redis.IntCmd
	RPush(string, ...interface{}) *redis.IntCmd
	BLPop(time.Duration, ...string) *redis.StringSliceCmd
	LPop(string) *redis.StringCmd
}
