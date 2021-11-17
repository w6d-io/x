package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) CreateIndexes(opt mongo.IndexModel) error {
	log := logx.WithName(nil, "Create Index")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	// opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	err := i.Client.CreateOne(ctx, i.Database, i.Collection, opt)
	if err != nil {
		log.Error(err, "create index err")
	}
	return err
}

func (c *Client) CreateOne(ctx context.Context, database string, collection string, opt mongo.IndexModel) error {
	col := c.client.Database(database).Collection(collection)

	_, err := col.Indexes().CreateOne(ctx, opt)

	return err
}
