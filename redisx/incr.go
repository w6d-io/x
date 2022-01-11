package redisx

import (
	"context"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Incr ...
func (r *RedisDB) Incr(ctx context.Context, key string) (int64, error) {
	log := logx.WithName(ctx, "Redis.Incr")
	if err := r.Connect(); err != nil {
		return 0, errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.Incr(key).Result()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "failed", "key", key)
		return 0, err
	} else if err != nil && err.Error() == NilRedis {
		log.Error(err, "didn't find the key", "key", key)
		return 0, err
	}
	return result, nil
}
