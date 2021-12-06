package kafkax

import (
	"time"
)

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	Protocol           string
	Mechanisms         string
	Async              bool
	SessionTimeout     time.Duration
	MaxPollInterval    time.Duration
	WriteTimeout       time.Duration
	ReadTimeout        time.Duration
	BatchTimeout       time.Duration
	MaxWait            time.Duration
	StatInterval       time.Duration
	NumPartitions      int
	ReplicationFactor  int
	MinBytes           int
	MaxBytes           int
	AuthKafka          bool
	FullStats          bool
	Debugs             []string
	GroupInstanceID    string
	ConfigMapKey       string
	AutoOffsetReset    string
	EnableRebalance    bool
	EnablePartitionEOF bool
	Headers            []Header
}

// NewOptions ...
func NewOptions(opts ...Option) Options {
	opt := Options{
		Protocol:           "SASL_SSL",
		Mechanisms:         "PLAIN",
		Async:              true,
		SessionTimeout:     10 * time.Second,
		MaxPollInterval:    5 * time.Minute,
		WriteTimeout:       10 * time.Second,
		ReadTimeout:        10 * time.Second,
		BatchTimeout:       1 * time.Millisecond,
		MaxWait:            2 * time.Millisecond,
		StatInterval:       5 * time.Second,
		MinBytes:           10e3,
		MaxBytes:           10e6,
		NumPartitions:      1,
		ReplicationFactor:  3,
		AuthKafka:          false,
		FullStats:          false,
		GroupInstanceID:    "",
		Debugs:             []string{},
		ConfigMapKey:       "kafka",
		AutoOffsetReset:    "earliest",
		EnableRebalance:    false,
		EnablePartitionEOF: false,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Protocol option
func Protocol(p string) Option {
	return func(o *Options) {
		o.Protocol = p
	}
}

// Mechanisms option
func Mechanisms(m string) Option {
	return func(o *Options) {
		o.Mechanisms = m
	}
}

// Async option
func Async(b bool) Option {
	return func(o *Options) {
		o.Async = b
	}
}

// WriteTimeout option
func WriteTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = t
	}
}

// MaxWait option
func MaxWait(t time.Duration) Option {
	return func(o *Options) {
		o.MaxWait = t
	}
}

// StatInterval option
func StatInterval(t time.Duration) Option {
	return func(o *Options) {
		o.StatInterval = t
	}
}

// MaxBytes option
func MaxBytes(b int) Option {
	return func(o *Options) {
		o.MaxBytes = b
	}
}

// MinBytes option
func MinBytes(b int) Option {
	return func(o *Options) {
		o.MinBytes = b
	}
}

// NumPartitions option
func NumPartitions(n int) Option {
	return func(o *Options) {
		o.NumPartitions = n
	}
}

// ReplicationFactor option
func ReplicationFactor(r int) Option {
	return func(o *Options) {
		o.ReplicationFactor = r
	}
}

// AuthKafka option
func AuthKafka(b bool) Option {
	return func(o *Options) {
		o.AuthKafka = b
	}
}

// FullStats option
func FullStats(b bool) Option {
	return func(o *Options) {
		o.FullStats = b
	}
}

// Debugs option
func Debugs(d []string) Option {
	return func(o *Options) {
		o.Debugs = d
	}
}

// SessionTimeout option
func SessionTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.SessionTimeout = t
	}
}

// MaxPollInterval option
func MaxPollInterval(t time.Duration) Option {
	return func(o *Options) {
		o.MaxPollInterval = t
	}
}

// GroupInstanceID option
func GroupInstanceID(s string) Option {
	return func(o *Options) {
		o.GroupInstanceID = s
	}
}

// ConfigMapKey option
func ConfigMapKey(s string) Option {
	return func(o *Options) {
		o.ConfigMapKey = s
	}
}

// Rebalance option
func Rebalance(enable bool) Option {
	return func(o *Options) {
		o.EnableRebalance = enable
	}
}

// PartitionEOF option
func PartitionEOF(enable bool) Option {
	return func(o *Options) {
		o.EnablePartitionEOF = enable
	}
}

// EarliestOffset Strategy
func EarliestOffset() Option {
	return func(o *Options) {
		o.AutoOffsetReset = "earliest"
	}
}

// WithHeaders ...
func WithHeaders(headers []Header) Option {
	return func(o *Options) {
		o.Headers = headers
	}
}
