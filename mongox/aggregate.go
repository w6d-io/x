package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// Aggregate operator over mongo client
func (m *MongoDB) Aggregate(pipeline mongo.Pipeline, data interface{}) error {
	log := logx.WithName(context.TODO(), "Aggregate")
	ctx, cancel := context.WithTimeout(context.Background(), m.options.Timeout)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	cursor, err := m.GetCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Error(err, "find")
		return err
	}
	if err := m.SetCursor(cursor).All(ctx, data); err != nil {
		log.Error(err, "get data")
		return err
	}

	log.WithValues("data", data).V(GetLogLevel(data)).Info("result from aggregate")

	return nil
}
