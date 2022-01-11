package redisx

import (
	"context"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// HDel ...
func (r *RedisDB) HDel(ctx context.Context, key string, fields ...string) error {
	log := logx.WithName(ctx, "HDel")
	if err := r.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	err := r.ClientAPI.HDel(key, fields...).Err()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "Redis.HDel Failed", "key", key, "fields", fields)
		return err
	} else if err != nil && err.Error() == NilRedis {
		log.Error(err, "Redis.HDel Nothing to delete", "key", key, "fields", fields)
		return err
	}
	return nil
}
