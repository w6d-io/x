package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockIndex is the internal mock index
type MockIndex struct {
	CursorAPI
	IndexResult string
	ErrIndex    error
}

// CreateOne is an internal mock method
func (i *MockIndex) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return i.IndexResult, i.ErrIndex
}
