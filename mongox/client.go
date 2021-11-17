package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// GetClient return the mongo client recorded or create a new instance
func GetClient(opts ...*mgoOtions.ClientOptions) (*Client, error) {
	if db != nil && len(opts) == 0 {
		return db, nil
	}
	log := logx.WithName(nil, "GetClient")
	clt, err := mongo.NewClient(opts...)
	if err != nil {
		log.Error(err, "init mongo client failed")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = clt.Connect(ctx)
	if err != nil {
		log.Error(err, "init mongo client failed")
		return nil, err
	}
	db = &Client{
		client: clt,
		opts:   opts,
	}
	return db, nil
}

func (c *Client) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if c.client == nil {
		client, err := mongo.NewClient(c.opts...)
		if err != nil {
			return errorx.Wrap(err, "fail to create new mongo client")
		}
		c.client = client
		if err := c.client.Connect(ctx); err != nil {
			logx.WithName(ctx, "MongoDB").Error(err, "init mongo client failed")
			return errorx.Wrap(err, "mongo connection failed")
		}
	}

	if err := c.client.Ping(ctx, readpref.Primary()); err != nil {
		return errorx.Wrap(err, "ping failed")
	}
	return nil
}
