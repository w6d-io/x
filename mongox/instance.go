package mongox

import (
	"errors"

	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
)

// New return a mongo instance
func (cfg *Mongo) New(collection string, opts ...Option) (IInstance, error) {
	options := NewOptions(opts...)
	authSource := cfg.Name
	if cfg.AuthSource != "" {
		authSource = cfg.AuthSource
	}
	instance := &Instance{}

	clientOptions := mgoOtions.Client().SetAuth(
		mgoOtions.Credential{
			Username:   cfg.Username,
			Password:   cfg.Password,
			AuthSource: authSource,
		}).ApplyURI(cfg.URL)

	if options.ProtoCodec {
		clientOptions.SetRegistry(ProtoCodecRegistry().Build())
	}
	c, err := GetClient(clientOptions)

	if c == nil {
		return nil, errors.New("GetClient return nil")
	}
	instance.Client = c
	instance.Database = cfg.Name
	instance.Collection = collection
	return instance, err
}
