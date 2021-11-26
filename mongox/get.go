package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (m *MongoDB) Get(filter interface{}) (CursorAPI, error) {
	log := logx.WithName(nil, "Get")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}
	findOptions := options.Find()
	cursor, err := m.GetCollection().Find(ctx, filter, findOptions)
	if err != nil {
		log.Error(err, "find")
		return nil, err
	}

	return m.SetCursor(cursor), nil
}
