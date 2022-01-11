package redisx

import (
	"time"

	"github.com/go-redis/redis/v7"
)

// MockClient is the internal mock client
type MockClient struct {
	ClientAPI
	PingErr      error
	HDelErr      error
	GetValue     string
	GetErr       error
	HGetValue    string
	HGetErr      error
	HGetAllValue map[string]string
	HGetAllErr   error
	KeysValue    []string
	KeysErr      error
	ScanValue    []string
	ScanCursor   uint64
	ScanErr      error
	IncrValue    int64
	IncrErr      error
	SetValue     string
	SetErr       error
	HSetValue    int64
	HSetErr      error
}

// Ping is an internal mock method
func (r *MockClient) Ping() *redis.StatusCmd {
	return redis.NewStatusResult(
		"",
		r.PingErr,
	)
}

// Close is an internal mock method
func (r *MockClient) Close() error {
	return nil
}

// HDel is an internal mock method
func (r *MockClient) HDel(key string, fields ...string) *redis.IntCmd {
	return redis.NewIntResult(
		0,
		r.HDelErr,
	)
}

// Get is an internal mock method
func (r *MockClient) Get(key string) *redis.StringCmd {
	return redis.NewStringResult(
		r.GetValue,
		r.GetErr,
	)
}

// HGet is an internal mock method
func (r *MockClient) HGet(key string, field string) *redis.StringCmd {
	return redis.NewStringResult(
		r.HGetValue,
		r.HGetErr,
	)
}

// HGetAll is an internal mock method
func (r *MockClient) HGetAll(key string) *redis.StringStringMapCmd {
	return redis.NewStringStringMapResult(
		r.HGetAllValue,
		r.HGetAllErr,
	)
}

// Keys is an internal mock method
func (r *MockClient) Keys(pattern string) *redis.StringSliceCmd {
	return redis.NewStringSliceResult(
		r.KeysValue,
		r.KeysErr,
	)
}

// Scan is an internal mock method
func (r *MockClient) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return redis.NewScanCmdResult(
		r.ScanValue,
		r.ScanCursor,
		r.ScanErr,
	)
}

// Incr is an internal mock method
func (r *MockClient) Incr(key string) *redis.IntCmd {
	return redis.NewIntResult(
		r.IncrValue,
		r.IncrErr,
	)
}

// Set is an internal mock method
func (r *MockClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult(
		r.SetValue,
		r.SetErr,
	)
}

// HSet is an internal mock method
func (r *MockClient) HSet(key string, values ...interface{}) *redis.IntCmd {
	return redis.NewIntResult(
		r.HSetValue,
		r.HSetErr,
	)
}
