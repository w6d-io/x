package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/w6d-io/x/logx"
)

// SetCollection set collection name for the internal client
func (c *Client) SetCollection(collection string) {
	c.Collection = collection
}

// GetCollection return a CollectionAPI
func (c *Client) GetCollection() CollectionAPI {
	return &ClientCollection{
		c.Client.Database(c.Database).Collection(c.Collection),
	}
}

// SetCursor return a CursorAPI
func (c *Client) SetCursor(cursor *mongo.Cursor) CursorAPI {
	return &ClientCursor{
		cursor,
	}
}

// SetSingleResult return a SingleResultAPI
func (c *Client) SetSingleResult(singleresult *mongo.SingleResult) SingleResultAPI {
	return &ClientSingleResult{
		singleresult,
	}
}

// GetClient return the mongo client recorded or create a new instance
func GetClient(ctx context.Context, m *MongoDB) (ClientAPI, error) {
	log := logx.WithName(ctx, "GetClient")

	authSource := m.cfg.Name
	if m.cfg.AuthSource != "" {
		authSource = m.cfg.AuthSource
	}
	clientOptions := mgoOtions.Client().SetAuth(
		mgoOtions.Credential{
			Username:   m.cfg.Username,
			Password:   m.cfg.Password,
			AuthSource: authSource,
		}).ApplyURI(m.cfg.URL)

	if m.options.ProtoCodec {
		clientOptions.SetRegistry(ProtoCodecRegistry().Build())
	}

	clt, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Error(err, "init mongo client failed")
		return nil, err
	}
	return &Client{
		Client:     clt,
		Database:   m.cfg.Name,
		Collection: m.Collection,
	}, nil
}

// New return a mongo instance
func (cfg *Mongo) New(opts ...Option) MongoAPI {
	options := NewOptions(opts...)

	return &MongoDB{
		cfg:         cfg,
		options:     &options,
		isConnected: false,
	}
}

// SetCollection set collection name for the public client
func (m *MongoDB) SetCollection(collection string) MongoAPI {
	m.Collection = collection
	if m.ClientAPI != nil {
		m.ClientAPI.SetCollection(collection)
	}
	return m
}

// SetOptions set options for the public client
func (m *MongoDB) SetOptions(opts ...Option) MongoAPI {
	options := NewOptions(opts...)
	m.options = &options
	return m
}

// Connect public client to the internal client
func (m *MongoDB) Connect() error {
	log := logx.WithName(context.TODO(), "Connect")
	ctx, cancel := context.WithTimeout(context.Background(), m.options.Timeout)
	defer cancel()
	if m.ClientAPI == nil {
		c, err := GetClient(ctx, m)
		if err != nil {
			log.Error(err, "create mongo client failed")
			return err
		}
		m.ClientAPI = c
	}
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
