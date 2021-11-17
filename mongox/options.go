package mongox

// Option
type Option func(*Options)

// Options ...
type Options struct {
	ProtoCodec bool
}

// NewOptions ...
func NewOptions(opts ...Option) Options {
	opt := Options{
		ProtoCodec: false,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Protobuf codec option
func WithProtoCodec() Option {
	return func(o *Options) {
		o.ProtoCodec = true
	}
}
