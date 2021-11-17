package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) Delete(filter interface{}) error {
	log := logx.WithName(nil, "Delete")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := i.Client.DeleteOne(ctx, i.Database, i.Collection, filter)
	if err != nil {
		log.Error(err, "delete")
	}
	log.WithValues("delete_count", result).V(1).Info("delete result")
	return err

}

func (i *Instance) DeleteAll() error {
	log := logx.WithName(nil, "Delete All")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := i.Client.DeleteMany(ctx, i.Database, i.Collection)
	if err != nil {
		log.Error(err, "delete all")
	}
	log.WithValues("delete_count", result).V(1).Info("delete result")
	return err
}

func (c *Client) DeleteOne(ctx context.Context, database string, collection string, filter interface{}) (int64, error) {
	col := c.client.Database(database).Collection(collection)

	result, err := col.DeleteOne(ctx, filter)

	if err != nil {
		return -1, err
	}

	return result.DeletedCount, nil
}

func (c *Client) DeleteMany(ctx context.Context, database string, collection string) (int64, error) {
	col := c.client.Database(database).Collection(collection)

	filter := bson.M{}

	result, err := col.DeleteMany(ctx, filter)

	if err != nil {
		return -1, err
	}

	return result.DeletedCount, nil
}
