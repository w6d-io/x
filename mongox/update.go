package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Update single document from a collection based on filter input
func (m *MongoDB) Update(filter interface{}, update interface{}) error {
	log := logx.WithName(context.TODO(), "Update")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := m.GetCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error(err, "update")
		return err
	}
	log.WithValues("Update", result.ModifiedCount).V(1).Info("update result")
	return err
}

// Upsert single document from a collection based on filter input
func (m *MongoDB) Upsert(filter interface{}, update interface{}) error {
	log := logx.WithName(context.TODO(), "Upsert")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := m.GetCollection().ReplaceOne(ctx, filter, update)
	if err != nil {
		log.Error(err, "Upsert")
		return err
	}
	log.WithValues("Upsert", result.ModifiedCount).V(1).Info("upsert result")
	return err
}

// FindAndUpdate single document from a collection based on filter input
func (m *MongoDB) FindAndUpdate(filter interface{}, update interface{}, data interface{}) error {
	log := logx.WithName(context.TODO(), "FindOneAndUpdate")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	sr := m.GetCollection().FindOneAndUpdate(ctx, filter, update, &opt)

	if err := m.SetSingleResult(sr).Decode(data); err != nil {
		log.Error(err, "get data")
		return err
	}
	log.WithValues("data", data).V(1).Info("result from FindOneAndUpdate")
	return nil
}
