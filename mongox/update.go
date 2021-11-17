package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) Update(filter interface{}, update interface{}) error {
	log := logx.WithName(nil, "Update")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := i.Client.Update(ctx, i.Database, i.Collection, filter, update)
	if err != nil {
		log.Error(err, "update")
	}
	log.WithValues("Update", result).V(1).Info("update result")
	return err
}

func (i *Instance) Upsert(filter interface{}, update interface{}) error {
	log := logx.WithName(nil, "Upsert")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := i.Client.Upsert(ctx, i.Database, i.Collection, filter, update)
	if err != nil {
		log.Error(err, "Upsert")
	}
	log.WithValues("Upsert", result).V(1).Info("upsert result")
	return err
}

func (i *Instance) FindAndUpdate(filter interface{}, update interface{}, data interface{}) error {
	log := logx.WithName(nil, "FindOneAndUpdate")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	sr, err := i.Client.FindOneAndUpdate(ctx, i.Database, i.Collection, filter, update)
	if err != nil {
		log.Error(err, "Find one and update")
	}
	if err := sr.Decode(data); err != nil {
		log.Error(err, "get data")
		return err
	}
	log.WithValues("data", data).V(1).Info("result from FindOneAndUpdate")
	return nil
}

func (c *Client) Update(ctx context.Context, database string, collection string, filter interface{}, update interface{}) (int64, error) {
	col := c.client.Database(database).Collection(collection)

	result, err := col.UpdateOne(ctx, filter, update)

	if err != nil {
		return -1, err
	}
	return result.ModifiedCount, nil
}

func (c *Client) Upsert(ctx context.Context, database string, collection string, filter interface{}, update interface{}) (interface{}, error) {
	col := c.client.Database(database).Collection(collection)

	result, err := col.ReplaceOne(ctx, filter, update)

	if err != nil {
		return -1, err
	}
	return result.UpsertedID, nil
}

func (c *Client) FindOneAndUpdate(ctx context.Context, database string, collection string, filter interface{}, update interface{}) (ISingleResult, error) {
	col := c.client.Database(database).Collection(collection)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := col.FindOneAndUpdate(ctx, filter, update, &opt)

	if result.Err() != nil {
		return nil, result.Err()
	}
	return &SingleResult{
		singleResult: result,
	}, nil
}
