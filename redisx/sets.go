package redisx

import (
	"context"
	"encoding/json"
	"time"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Set set a key/value with an expiry duration, after this duration the key/value is deleted from redis
func (r *RedisDB) Set(ctx context.Context, key string, expiry time.Duration, data interface{}) error {
	log := logx.WithName(ctx, "Redis.Set")
	if err := r.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	err := r.ClientAPI.Set(key, data, expiry).Err()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "set failed", "key", key, "expiry", expiry)
		return err
	}

	return nil
}

// HSet ...
func (r *RedisDB) HSet(ctx context.Context, key string, field string, data interface{}) error {
	log := logx.WithName(ctx, "Redis.Hset")
	if err := r.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	eventBytes, err := json.Marshal(data)
	if err != nil {
		log.Error(err, "error marshaling data")
		return err
	}

	err = r.ClientAPI.HSet(key, field, eventBytes).Err()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "HSet Failed", "key", key, "field", field)
		return err
	}

	return nil
}
