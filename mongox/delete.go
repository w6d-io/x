package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Delete remove documents from collection based on input filter setting
func (m *MongoDB) Delete(filter interface{}) error {
	log := logx.WithName(nil, "Delete")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	result, err := m.GetCollection().DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "delete")
		return err
	}
	log.WithValues("delete_count", result.DeletedCount).V(1).Info("delete result")
	return nil

}

// DeleteAll remove all documents inside collection
func (m *MongoDB) DeleteAll() error {
	log := logx.WithName(nil, "Delete All")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	filter := bson.M{}
	result, err := m.GetCollection().DeleteMany(ctx, filter)
	if err != nil {
		log.Error(err, "delete all")
		return err
	}
	log.WithValues("delete_count", result.DeletedCount).V(1).Info("delete result")
	return nil
}
