package redisx

import (
	"context"
	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
	"time"
)

// RPush ...
func (r *RedisDB) RPush(ctx context.Context, key string, values ...interface{}) error {
	log := logx.WithName(ctx, "Redis.RPush")
	if err := r.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	if _, err := r.ClientAPI.RPush(key, values...).Result(); err != nil {
		log.Error(err, "fail to push", "key", key)
		return err
	}
	return nil
}

// BLPop ...
func (r *RedisDB) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	log := logx.WithName(ctx, "Redis.BLPop")
	if err := r.Connect(); err != nil {
		return []string{}, errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.BLPop(timeout, keys...).Result()
	if err != nil {
		log.Error(err, "fail to pop", "keys", keys)
		return []string{}, err
	}
	return result, nil
}

// LPop ...
func (r *RedisDB) LPop(ctx context.Context, key string) (string, error) {
	log := logx.WithName(ctx, "Redis.LPop")
	if err := r.Connect(); err != nil {
		return "", errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.LPop(key).Result()
	if err != nil {
		log.Error(err, "fail to pop", "key", key)
		return "", err
	}
	return result, nil
}
