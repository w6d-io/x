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
	CreateManyIndexResult   []string
	ListSpecificationResult []*mongo.IndexSpecification
	DropOneResult           bson.Raw
	ErrCreateIndex          error
	ErrCreateManyIndex      error
	ErrListSpecifications   error
	ErrDropOne              error
}

// CreateOne is an internal mock method
func (i *MockIndex) CreateOne(_ context.Context, _ mongo.IndexModel, _ ...*options.CreateIndexesOptions) (string, error) {
	return i.CreateIndexResult, i.ErrCreateIndex
}

// CreateMany is an internal mock method
func (i *MockIndex) CreateMany(_ context.Context, _ []mongo.IndexModel, _ ...*options.CreateIndexesOptions) ([]string, error) {
	return i.CreateManyIndexResult, i.ErrCreateManyIndex
}

// ListSpecifications is an internal mock method
func (i *MockIndex) ListSpecifications(_ context.Context, _ ...*options.ListIndexesOptions) ([]*mongo.IndexSpecification, error) {
	return i.ListSpecificationResult, i.ErrListSpecifications
}

// DropOne is an internal mock method
func (i *MockIndex) DropOne(_ context.Context, _ string, _ ...*options.DropIndexesOptions) (bson.Raw, error) {
	return i.DropOneResult, i.ErrDropOne
}
