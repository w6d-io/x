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
	RPushValue   int64
	RPushErr     error
	BLPopValue   []string
	BLPopErr     error
	LPopValue    string
	LPopErr      error
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
func (r *MockClient) HDel(_ string, _ ...string) *redis.IntCmd {
	return redis.NewIntResult(
		0,
		r.HDelErr,
	)
}

// Get is an internal mock method
func (r *MockClient) Get(_ string) *redis.StringCmd {
	return redis.NewStringResult(
		r.GetValue,
		r.GetErr,
	)
}

// HGet is an internal mock method
func (r *MockClient) HGet(_ string, _ string) *redis.StringCmd {
	return redis.NewStringResult(
		r.HGetValue,
		r.HGetErr,
	)
}

// HGetAll is an internal mock method
func (r *MockClient) HGetAll(_ string) *redis.StringStringMapCmd {
	return redis.NewStringStringMapResult(
		r.HGetAllValue,
		r.HGetAllErr,
	)
}

// Keys is an internal mock method
func (r *MockClient) Keys(_ string) *redis.StringSliceCmd {
	return redis.NewStringSliceResult(
		r.KeysValue,
		r.KeysErr,
	)
}

// Scan is an internal mock method
func (r *MockClient) Scan(_ uint64, _ string, _ int64) *redis.ScanCmd {
	return redis.NewScanCmdResult(
		r.ScanValue,
		r.ScanCursor,
		r.ScanErr,
	)
}

// Incr is an internal mock method
func (r *MockClient) Incr(_ string) *redis.IntCmd {
	return redis.NewIntResult(
		r.IncrValue,
		r.IncrErr,
	)
}

// Set is an internal mock method
func (r *MockClient) Set(_ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	return redis.NewStatusResult(
		r.SetValue,
		r.SetErr,
	)
}

// HSet is an internal mock method
func (r *MockClient) HSet(_ string, _ ...interface{}) *redis.IntCmd {
	return redis.NewIntResult(
		r.HSetValue,
		r.HSetErr,
	)
}

// RPush is an internal mock method
func (r *MockClient) RPush(_ string, _ ...interface{}) *redis.IntCmd {
	return redis.NewIntResult(
		r.RPushValue,
		r.RPushErr,
	)
}

// BLPop is an internal mock method
func (r *MockClient) BLPop(_ time.Duration, _ ...string) *redis.StringSliceCmd {
	return redis.NewStringSliceResult(
		r.BLPopValue,
		r.BLPopErr,
	)
}

// LPop is an internal mock method
func (r *MockClient) LPop(_ string) *redis.StringCmd {
	return redis.NewStringResult(
		r.LPopValue,
		r.LPopErr,
	)
}
