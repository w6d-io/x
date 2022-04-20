package mongox

import "time"

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	ProtoCodec bool
	Timeout    time.Duration
}

// NewOptions ...
func NewOptions(opts ...Option) Options {
	opt := Options{
		ProtoCodec: false,
		Timeout:    10 * time.Second,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// WithProtoCodec option
func WithProtoCodec() Option {
	return func(o *Options) {
		o.ProtoCodec = true
	}
}

// Timeout option
func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}
