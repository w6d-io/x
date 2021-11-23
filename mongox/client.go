package mongox

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/w6d-io/x/logx"
)

func (c *Client) GetCollection() CollectionAPI {
	return &ClientCollection{
		c.Client.Database(c.Database).Collection(c.Collection),
	}
}

func (c *Client) SetCursor(cursor *mongo.Cursor) CursorAPI {
	return &ClientCursor{
		cursor,
	}
}

func (c *Client) SetSingleResult(singleresult *mongo.SingleResult) SingleResultAPI {
	return &ClientSingleResult{
		singleresult,
	}
}

// GetClient return the mongo client recorded or create a new instance
func GetClient(database string, collection string, opts ...*mgoOtions.ClientOptions) (ClientAPI, error) {
	log := logx.WithName(nil, "GetClient")
	clt, err := mongo.NewClient(opts...)
	if err != nil {
		log.Error(err, "init mongo client failed")
		return nil, err
	}
	return &Client{
		Client:     clt,
		opts:       opts,
		Database:   database,
		Collection: collection,
	}, nil
}

// New return a mongo instance
func (cfg *Mongo) New(collection string, opts ...Option) (MongoAPI, error) {
	options := NewOptions(opts...)
	authSource := cfg.Name
	if cfg.AuthSource != "" {
		authSource = cfg.AuthSource
	}
	clientOptions := mgoOtions.Client().SetAuth(
		mgoOtions.Credential{
			Username:   cfg.Username,
			Password:   cfg.Password,
			AuthSource: authSource,
		}).ApplyURI(cfg.URL)

	if options.ProtoCodec {
		clientOptions.SetRegistry(ProtoCodecRegistry().Build())
	}
	c, err := GetClient(cfg.Name, collection, clientOptions)

	if err != nil {
		return nil, errors.New("GetClient return nil")
	}
	return &MongoDB{
		ClientAPI:   c,
		isConnected: false,
		Database:    cfg.Name,
		Collection:  collection,
	}, nil
}

func (m *MongoDB) Connect() error {
	log := logx.WithName(nil, "Connect")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if !m.isConnected {
		err := m.ClientAPI.Connect(ctx)
		if err != nil {
			log.Error(err, "init mongo client failed")
			return err
		}
		m.isConnected = true
	}
	if err := m.ClientAPI.Ping(ctx, readpref.Primary()); err != nil {
		log.Error(err, "ping failed")
		return err
	}
	return nil
}
