package mongox

import "time"

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	ProtoCodec bool
	StrCodec   bool
	Timeout    time.Duration
}

// NewOptions ...
func NewOptions(opts ...Option) Options {
	opt := Options{
		ProtoCodec: false,
		StrCodec:   false,
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

// WithStrCodec option
func WithStrCodec() Option {
	return func(o *Options) {
		o.StrCodec = true
	}
}

// Timeout option
func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}
