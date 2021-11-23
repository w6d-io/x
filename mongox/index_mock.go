package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockIndex struct {
	CursorAPI
	IndexResult string
	ErrIndex    error
}

func (i *MockIndex) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return i.IndexResult, i.ErrIndex
}
