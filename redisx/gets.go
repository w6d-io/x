package redisx

import (
	"context"
	"encoding/json"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Get ...
func (r *RedisDB) Get(ctx context.Context, key string) (string, error) {
	log := logx.WithName(ctx, "Redis.Get")
	if err := r.Connect(); err != nil {
		return "", errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.Get(key).Result()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "Redis.Get Failed", "key", key)
		return "", err
	} else if err != nil && err.Error() == NilRedis {
		return "", err
	}

	return result, nil
}

// HGet ...
func (r *RedisDB) HGet(ctx context.Context, key string, field string, response interface{}) error {
	log := logx.WithName(ctx, "Redis.HGet")
	if err := r.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.HGet(key, field).Result()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "Redis.HGet Failed. Key", "key", key)
		return err
	} else if err != nil && err.Error() == NilRedis {
		return err
	}

	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		return err
	}

	return nil
}

// HGetAll ...
func (r *RedisDB) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	log := logx.WithName(ctx, "Redis.HGetAll")
	if err := r.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.HGetAll(key).Result()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "Redis.HGetAll Failed", "key", key)
		return nil, err
	} else if err != nil && err.Error() == NilRedis {
		return nil, err
	}

	return result, nil
}

// Keys return all keys from redis if the pattern is set to "*", not safe, use Scan instead
func (r *RedisDB) Keys(ctx context.Context, pattern string) ([]string, error) {
	log := logx.WithName(ctx, "Redis.Key")
	if err := r.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}
	result, err := r.ClientAPI.Keys(pattern).Result()
	if err != nil && err.Error() != NilRedis {
		log.Error(err, "Redis.Keys Failed", "pattern", pattern)
		return nil, err
	} else if err != nil && err.Error() == NilRedis {
		return nil, err
	}
	return result, nil
}

// Scan returns every keys from a redis database safely
func (r *RedisDB) Scan(ctx context.Context, match string, count int) (map[string]string, error) {
	log := logx.WithName(ctx, "Redis.Scan").WithValues(logx.GetLogValues(ctx)...)
	if err := r.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}
	var cursor uint64
	var err error
	var keys []string
	keysMap := make(map[string]string, 0)

	for {
		keys, cursor, err = r.ClientAPI.Scan(cursor, match, int64(count)).Result()
		if err != nil {
			log.Error(err, "error when i try to scan")
			return nil, err
		}
		for _, i := range keys {
			keysMap[i] = ""
		}
		if cursor == 0 {
			return keysMap, nil
		}
	}
}
