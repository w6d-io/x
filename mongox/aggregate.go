package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (m *MongoDB) Aggregate(pipeline mongo.Pipeline) (CursorAPI, error) {
	log := logx.WithName(nil, "Aggregate")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}

	cursor, err := m.GetCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Error(err, "find")
		return nil, err
	}

	return m.SetCursor(cursor), nil
}
