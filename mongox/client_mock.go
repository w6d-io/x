package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MockClient struct {
	ClientAPI
	ErrConnect   error
	ErrPing      error
	ErrInsertOne error
	ErrBulkWrite error
	// NumSessionsInProgress   int
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
	IndexResult             string
	ErrIndex                error
}

func (p *MockClient) Connect(ctx context.Context) error {
	return p.ErrConnect
}

func (p *MockClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return p.ErrPing
}

// func (p *MockClient) NumberSessionsInProgress() int {
// 	return p.NumSessionsInProgress
// }

func (p *MockClient) GetCollection() CollectionAPI {
	return &MockCollection{
		ErrInsertOne:           p.ErrInsertOne,
		ErrBulkWrite:           p.ErrBulkWrite,
		InsertOneResult:        p.InsertOneResult,
		BulkWriteResult:        p.BulkWriteResult,
		ErrFind:                p.ErrFind,
		ErrDeleteOne:           p.ErrDeleteOne,
		ErrDeleteMany:          p.ErrDeleteMany,
		ErrUpdateOne:           p.ErrUpdateOne,
		ErrReplaceOne:          p.ErrReplaceOne,
		UpdateOneResult:        p.UpdateOneResult,
		ReplaceOneResult:       p.ReplaceOneResult,
		FindOneAndUpdateResult: p.FindOneAndUpdateResult,
		AggregateResult:        p.AggregateResult,
		ErrAggregate:           p.ErrAggregate,
		IndexResult:            p.IndexResult,
		ErrIndex:               p.ErrIndex,
	}
}

func (p *MockClient) SetCursor(*mongo.Cursor) CursorAPI {
	return &MockCursor{
		ErrorCursorAll: p.ErrorCursorAll,
	}
}

func (p *MockClient) SetSingleResult(*mongo.SingleResult) SingleResultAPI {
	return &MockSingleResult{
		ErrorSingleResultDecode: p.ErrorSingleResultDecode,
	}
}