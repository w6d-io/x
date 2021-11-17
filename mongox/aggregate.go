package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) Aggregate(pipeline mongo.Pipeline, data interface{}) error {
	log := logx.WithName(nil, "Aggregate")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	cursor, err := i.Client.Aggregate(ctx, i.Database, i.Collection, pipeline)
	if err != nil {
		log.Error(err, "find")
		return err
	}
	if err := cursor.All(ctx, data); err != nil {
		log.Error(err, "get data")
		return err
	}
	log.WithValues("data", data).V(1).Info("result from aggregate")
	return nil
}

func (c *Client) Aggregate(ctx context.Context, database string, collection string, pipeline mongo.Pipeline) (ICursor, error) {
	col := c.client.Database(database).Collection(collection)

	cur, err := col.Aggregate(
		ctx,
		pipeline,
	)

	if err != nil {
		return nil, err
	}

	return &Cursor{cursor: cur}, err
}
