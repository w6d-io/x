package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	CollectionAPI
	ErrInsertOne           error
	ErrBulkWrite           error
	ErrFind                error
	ErrDeleteOne           error
	ErrDeleteMany          error
	InsertOneResult        mongo.InsertOneResult
	BulkWriteResult        mongo.BulkWriteResult
	DeleteResult           mongo.DeleteResult
	ErrUpdateOne           error
	ErrReplaceOne          error
	UpdateOneResult        mongo.UpdateResult
	ReplaceOneResult       mongo.UpdateResult
	FindOneAndUpdateResult mongo.SingleResult
	AggregateResult        mongo.Cursor
	ErrAggregate           error
	IndexResult            string
	ErrIndex               error
}

func (c *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*mgoOtions.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return &c.InsertOneResult, c.ErrInsertOne
}

func (c *MockCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*mgoOtions.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return &c.BulkWriteResult, c.ErrBulkWrite
}

func (c *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*mgoOtions.FindOptions) (*mongo.Cursor, error) {
	return &mongo.Cursor{}, c.ErrFind
}

func (c *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &c.DeleteResult, c.ErrDeleteOne
}

func (c *MockCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return &c.DeleteResult, c.ErrDeleteMany
}

func (c *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return &c.UpdateOneResult, c.ErrUpdateOne
}

func (c *MockCollection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return &c.ReplaceOneResult, c.ErrReplaceOne
}

func (c *MockCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return &c.FindOneAndUpdateResult
}

func (c *MockCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return &c.AggregateResult, c.ErrAggregate
}

func (c *MockCollection) GetIndex() IndexAPI {
	return &MockIndex{
		IndexResult: c.IndexResult,
		ErrIndex:    c.ErrIndex,
	}
}
