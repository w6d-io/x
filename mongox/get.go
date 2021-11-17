package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) Get(filter interface{}, data interface{}) error {
	log := logx.WithName(nil, "Get")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	cursor, err := i.Client.Find(ctx, i.Database, i.Collection, filter)
	if err != nil {
		log.Error(err, "find")
		return err
	}

	if err := cursor.All(ctx, data); err != nil {
		log.Error(err, "get data")
		return err
	}
	log.WithValues("data", data).V(1).Info("result from search")
	return nil
}

func (c *Client) Find(ctx context.Context, database string, collection string, filter interface{}) (ICursor, error) {
	col := c.client.Database(database).Collection(collection)

	findOptions := options.Find()

	cur, err := col.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	return &Cursor{cursor: cur}, err
}
