package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (m *MongoDB) Insert(value interface{}) error {
	log := logx.WithName(nil, "Insert")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	result, err := m.GetCollection().InsertOne(ctx, value)
	if err != nil {
		log.Error(err, "insert")
		return err
	}
	if result != nil {
		log.WithValues("insert_id", result.InsertedID).V(1).Info("insert done")
	}
	log.V(1).Info("insert with no error")
	return nil

}

func (m *MongoDB) InsertBulk(operations []*mongo.UpdateOneModel) error {
	log := logx.WithName(nil, "Insert Bulk")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	models := make([]mongo.WriteModel, len(operations))

	for i, op := range operations {
		models[i] = op
	}

	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)

	_, err := m.GetCollection().BulkWrite(ctx, models, &bulkOption)
	if err != nil {
		log.Error(err, "bulk")
		return err
	}
	log.V(1).Info("Bulk with no error")
	return err
}
