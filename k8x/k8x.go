package k8x

import (
	"bytes"
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	Namespace  string
	Context    string
	ConfigPath string
	Timeout    string
}

// Helper is K8X methods interface
type Helper interface {
	GetClient(context.Context) (client.Client, error)
	GetConfig(context.Context) (*rest.Config, error)
	SetClient(client.Client)
	SetConfig(config *rest.Config)
}

type k8x struct {
	configPath string
	context    string
	namespace  string
	timeout    string
	client     client.Client
	config     *rest.Config
}

// NewOptions ...
func NewOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Namespace return the namespace option setting
func Namespace(n string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

// Timeout return the timeout option setting
func Timeout(n string) Option {
	return func(o *Options) {
		o.Timeout = n
	}
}

// Context return the context option setting
func Context(c string) Option {
	return func(o *Options) {
		o.Context = c
	}
}

// ConfigPath return the config path option setting
func ConfigPath(c string) Option {
	return func(o *Options) {
		o.ConfigPath = c
	}
}

// New return an empty k8s configuration instance
func New(opts ...Option) Helper {
	options := NewOptions(opts...)
	return &k8x{
		configPath: options.ConfigPath,
		context:    options.Context,
		namespace:  options.Namespace,
		timeout:    options.Timeout,
	}
}

// GetClient return the client.Client or an error
func (k *k8x) GetClient(ctx context.Context) (client.Client, error) {
	log := logx.WithName(ctx, "GetClient")
	if k.client != nil {
		return k.client, nil
	}
	config, err := k.GetConfig(ctx)
	if err != nil {
		return nil, err
	}
	cli, err := client.New(config, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		log.Error(err, "build kubernetes client failed")
		return nil, &errorx.Error{Cause: err, Message: "get namespace failed"}
	}
	k.client = cli
	return k.client, nil
}

// GetConfig return the rest.Config or an error
func (k *k8x) GetConfig(ctx context.Context) (*rest.Config, error) {
	log := logx.WithName(ctx, "GetConfig")
	if k.config != nil {
		return k.config, nil
	}
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if k.configPath != "" {
		loadingRules.ExplicitPath = k.configPath
	}
	configOverrides := &clientcmd.ConfigOverrides{Timeout: k.timeout}
	if k.context != "" {
		configOverrides.CurrentContext = k.context
	}
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	if k.namespace == "" {
		namespace, _, err := config.Namespace()
		if err != nil {
			log.Error(err, "get namespace failed")
			return nil, &errorx.Error{Cause: err, Message: "get namespace failed"}
		}
		k.namespace = namespace
	}
	clientConfig, err := config.ClientConfig()
	if err == nil {
		k.config = clientConfig
	}
	return k.config, err
}

// SetClient records the client in the structure
func (k *k8x) SetClient(cli client.Client) {
	k.client = cli
}

// SetConfig records the config in the structure
func (k *k8x) SetConfig(config *rest.Config) {
	k.config = config
}

// GetObjectContain ...
func GetObjectContain(obj runtime.Object) string {
	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, nil, nil, json.SerializerOptions{Yaml: true})
	buf := new(bytes.Buffer)
	if err := s.Encode(obj, buf); obj == nil || err != nil {
		return "<ERROR>\n"
	}
	return buf.String()
}
