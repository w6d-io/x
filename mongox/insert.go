package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) Insert(value interface{}) error {
	log := logx.WithName(nil, "Insert")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	result, err := i.Client.InsertOne(ctx, i.Database, i.Collection, value)
	if err != nil {
		log.Error(err, "insert")
		return err
	}
	if result != nil {
		log.WithValues("insert_id", result).V(1).Info("insert done")
		return nil
	}
	log.V(1).Info("insert with no error")
	return err

}

func (i *Instance) InsertBulk(operations []*mongo.UpdateOneModel) error {
	log := logx.WithName(nil, "Insert Bulk")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	err := i.Client.InsertBulk(ctx, i.Database, i.Collection, operations)
	if err != nil {
		log.Error(err, "bulk")
		return err
	}
	log.V(1).Info("Bulk with no error")
	return err
}

func (c *Client) InsertOne(ctx context.Context, database string, collection string, event interface{}) (interface{}, error) {
	col := c.client.Database(database).Collection(collection)

	result, err := col.InsertOne(ctx, event)

	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (c *Client) InsertBulk(ctx context.Context, database string, collection string, operations []*mongo.UpdateOneModel) error {
	col := c.client.Database(database).Collection(collection)

	models := make([]mongo.WriteModel, len(operations))

	for i, op := range operations {
		models[i] = op
	}

	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)

	_, err := col.BulkWrite(
		ctx,
		models,
		&bulkOption,
	)

	return err
}
