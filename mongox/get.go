package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Get return documents from collection based on input filter
func (m *MongoDB) Get(filter interface{}, data interface{}) error {
	log := logx.WithName(nil, "Get")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	findOptions := options.Find()
	cursor, err := m.GetCollection().Find(ctx, filter, findOptions)
	if err != nil {
		log.Error(err, "find")
		return err
	}

	if err := m.SetCursor(cursor).All(ctx, data); err != nil {
		log.Error(err, "get data")
		return err
	}
	log.WithValues("data", data).V(1).Info("result from search")
	return nil
}
