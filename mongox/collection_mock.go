package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollection is the internal mock collection
type MockCollection struct {
	CollectionAPI
	ErrInsertOne            error
	ErrBulkWrite            error
	ErrFind                 error
	ErrDeleteOne            error
	ErrDeleteMany           error
	InsertOneResult         mongo.InsertOneResult
	BulkWriteResult         mongo.BulkWriteResult
	DeleteResult            mongo.DeleteResult
	ErrUpdateOne            error
	ErrReplaceOne           error
	UpdateOneResult         mongo.UpdateResult
	ReplaceOneResult        mongo.UpdateResult
	FindOneAndUpdateResult  mongo.SingleResult
	AggregateResult         mongo.Cursor
	ErrAggregate            error
	CreateIndexResult       string
	CreateManyIndexResult   []string
	ErrCreateIndex          error
	ErrCreateManyIndex      error
	ListSpecificationResult []*mongo.IndexSpecification
	ErrListSpecifications   error
	DropOneResult           bson.Raw
	ErrDropOne              error
}

// InsertOne is an internal mock method
func (c *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*mgoOtions.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &c.InsertOneResult, c.ErrInsertOne
}

// BulkWrite is an internal mock method
func (c *MockCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*mgoOtions.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return &c.BulkWriteResult, c.ErrBulkWrite
}

// Find is an internal mock method
func (c *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*mgoOtions.FindOptions) (*mongo.Cursor, error) {
	return &mongo.Cursor{}, c.ErrFind
}

// DeleteOne is an internal mock method
func (c *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &c.DeleteResult, c.ErrDeleteOne
}

// DeleteMany is an internal mock method
func (c *MockCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &c.DeleteResult, c.ErrDeleteMany
}

// UpdateOne is an internal mock method
func (c *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &c.UpdateOneResult, c.ErrUpdateOne
}

// ReplaceOne is an internal mock method
func (c *MockCollection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return &c.ReplaceOneResult, c.ErrReplaceOne
}

// FindOneAndUpdate is an internal mock method
func (c *MockCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return &c.FindOneAndUpdateResult
}

// Aggregate is an internal mock method
func (c *MockCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return &c.AggregateResult, c.ErrAggregate
}

// GetIndex is an internal mock method
func (c *MockCollection) GetIndex() IndexAPI {
	return &MockIndex{
		CreateIndexResult:       c.CreateIndexResult,
		CreateManyIndexResult:   c.CreateManyIndexResult,
		ErrCreateIndex:          c.ErrCreateIndex,
		ErrCreateManyIndex:      c.ErrCreateManyIndex,
		ListSpecificationResult: c.ListSpecificationResult,
		ErrListSpecifications:   c.ErrListSpecifications,
		DropOneResult:           c.DropOneResult,
		ErrDropOne:              c.ErrDropOne,
	}
}
