package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MockClient is the internal mock client
type MockClient struct {
	ClientAPI
	ErrConnect              error
	ErrPing                 error
	ErrInsertOne            error
	ErrBulkWrite            error
	InsertOneResult         mongo.InsertOneResult
	BulkWriteResult         mongo.BulkWriteResult
	ErrFind                 error
	ErrDeleteOne            error
	ErrDeleteMany           error
	ErrorCursorAll          error
	ErrUpdateOne            error
	ErrReplaceOne           error
	UpdateOneResult         mongo.UpdateResult
	ReplaceOneResult        mongo.UpdateResult
	FindOneAndUpdateResult  mongo.SingleResult
	ErrorSingleResultDecode error
	AggregateResult         mongo.Cursor
	ErrAggregate            error
	CreateIndexResult       string
	ErrCreateIndex          error
	ListSpecificationResult []*mongo.IndexSpecification
	ErrListSpecifications   error
	DropOneResult           bson.Raw
	ErrDropOne              error
}

// SetCollection is an internal mock method
func (p *MockClient) SetCollection(collection string) {
}

// Connect is an internal mock method
func (p *MockClient) Connect(ctx context.Context) error {
	return p.ErrConnect
}

// Ping is an internal mock method
func (p *MockClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return p.ErrPing
}

// GetCollection is an internal mock method
func (p *MockClient) GetCollection() CollectionAPI {
	return &MockCollection{
		ErrInsertOne:            p.ErrInsertOne,
		ErrBulkWrite:            p.ErrBulkWrite,
		InsertOneResult:         p.InsertOneResult,
		BulkWriteResult:         p.BulkWriteResult,
		ErrFind:                 p.ErrFind,
		ErrDeleteOne:            p.ErrDeleteOne,
		ErrDeleteMany:           p.ErrDeleteMany,
		ErrUpdateOne:            p.ErrUpdateOne,
		ErrReplaceOne:           p.ErrReplaceOne,
		UpdateOneResult:         p.UpdateOneResult,
		ReplaceOneResult:        p.ReplaceOneResult,
		FindOneAndUpdateResult:  p.FindOneAndUpdateResult,
		AggregateResult:         p.AggregateResult,
		ErrAggregate:            p.ErrAggregate,
		CreateIndexResult:       p.CreateIndexResult,
		ErrCreateIndex:          p.ErrCreateIndex,
		ListSpecificationResult: p.ListSpecificationResult,
		ErrListSpecifications:   p.ErrListSpecifications,
		DropOneResult:           p.DropOneResult,
		ErrDropOne:              p.ErrDropOne,
	}
}

// SetCursor is an internal mock method
func (p *MockClient) SetCursor(*mongo.Cursor) CursorAPI {
	return &MockCursor{
		ErrorCursorAll: p.ErrorCursorAll,
	}
}

// SetSingleResult is an internal mock method
func (p *MockClient) SetSingleResult(*mongo.SingleResult) SingleResultAPI {
	return &MockSingleResult{
		ErrorSingleResultDecode: p.ErrorSingleResultDecode,
	}
}
