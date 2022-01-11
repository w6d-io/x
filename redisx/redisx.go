package redisx

import (
	"github.com/go-redis/redis/v7"
)

// Redis connexion struct
type Redis struct {
	AddressSvc string `mapstructure:"address"`
	Port       string `mapstructure:"port"`
	Password   string `mapstructure:"password"`
	DB         int    `mapstructure:"db"`
}

// Client is the internal Redis Client
type Client struct {
	*redis.Client
}

// RedisDB is the public Redis Instance
type RedisDB struct {
	ClientAPI
	// From configuration section
	cfg *Redis
}
