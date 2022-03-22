package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockIndex is the internal mock index
type MockIndex struct {
	CursorAPI
	CreateIndexResult       string
	ListSpecificationResult []*mongo.IndexSpecification
	DropOneResult           bson.Raw
	ErrCreateIndex          error
	ErrListSpecifications   error
	ErrDropOne              error
}

// CreateOne is an internal mock method
func (i *MockIndex) CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return i.CreateIndexResult, i.ErrCreateIndex
}

// ListSpecifications is an internal mock method
func (i *MockIndex) ListSpecifications(ctx context.Context, opts ...*options.ListIndexesOptions) ([]*mongo.IndexSpecification, error) {
	return i.ListSpecificationResult, i.ErrListSpecifications
}

// DropOne is an internal mock method
func (i *MockIndex) DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return i.DropOneResult, i.ErrDropOne
}
