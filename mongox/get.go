package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Get return documents from collection based on input filter
func (m *MongoDB) Get(filter interface{}, data interface{}, findOptions ...*options.FindOptions) error {
	log := logx.WithName(context.TODO(), "Get")
	ctx, cancel := context.WithTimeout(context.Background(), m.options.Timeout)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	if findOptions == nil {
		findOptions = make([]*options.FindOptions, 1)
		findOptions[0] = options.Find()
	}

	cursor, err := m.GetCollection().Find(ctx, filter, findOptions...)
	if err != nil {
		log.Error(err, "find")
		return err
	}

	if err := m.SetCursor(cursor).All(ctx, data); err != nil {
		log.Error(err, "get data")
		return err
	}
	log.WithValues("data", data).V(GetLogLevel(data)).Info("result from search")
	return nil
}
