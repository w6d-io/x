package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (m *MongoDB) CountDocuments(filter interface{}, count *int64, countOptions ...*options.CountOptions) (err error) {
	log := logx.WithName(context.Background(), "CountDocuments")
	if count == nil {
		return errorx.New(nil, "count must not be null")
	}
	ctx, cancel := context.WithTimeout(context.Background(), m.options.Timeout)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}
	if countOptions == nil {
		countOptions = make([]*options.CountOptions, 1)
		countOptions[0] = options.Count()
	}

	*count, err = m.GetCollection().CountDocuments(ctx, filter, countOptions...)
	if err != nil {
		log.Error(err, "count documents")
		return err
	}
	return nil
}
