package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/w6d-io/x/logx"
)

// GetClient return the redis client recorded or create a new instance
func GetClient(ctx context.Context, r *RedisDB) ClientAPI {
	clt := redis.NewClient(&redis.Options{
		Addr:     r.cfg.AddressSvc + ":" + r.cfg.Port,
		Password: r.cfg.Password,
		DB:       r.cfg.DB,
	})
	return &Client{
		Client: clt,
	}
}

// New return a redis instance
func (cfg *Redis) New() RedisAPI {
	return &RedisDB{
		cfg: cfg,
	}
}

// Connect public client to the internal client
func (r *RedisDB) Connect() error {
	log := logx.WithName(context.TODO(), "Connect")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if r.ClientAPI == nil {
		r.ClientAPI = GetClient(ctx, r)
	}
	err := r.ClientAPI.Ping().Err()
	if err != nil {
		r.ClientAPI.Close()
		log.Error(err, "redis client is unreachable")
		return err
	}
	return nil
}
