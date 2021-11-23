package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoAPI interface {
	Connect() error
	Get(filter interface{}, data interface{}) error
	Insert(interface{}) error
	InsertBulk([]*mongo.UpdateOneModel) error
	Delete(filter interface{}) error
	DeleteAll() error
	Update(filter interface{}, update interface{}) error
	Upsert(filter interface{}, update interface{}) error
	FindAndUpdate(filter interface{}, update interface{}, data interface{}) error
	Aggregate(pipeline mongo.Pipeline, data interface{}) error
	CreateIndexes(mongo.IndexModel) error
	Incr(key string) (int64, error)
}

type ClientAPI interface {
	GetCollection() CollectionAPI
	SetCursor(*mongo.Cursor) CursorAPI
	SetSingleResult(*mongo.SingleResult) SingleResultAPI
	Connect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
}

type CollectionAPI interface {
	GetIndex() IndexAPI
	InsertOne(ctx context.Context, document interface{}, opts ...*mgoOtions.InsertOneOptions) (*mongo.InsertOneResult, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*mgoOtions.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
}

type CursorAPI interface {
	Next(ctx context.Context) bool
	All(ctx context.Context, results interface{}) error
	Decode(v interface{}) error
}

type SingleResultAPI interface {
	Decode(v interface{}) error
}

type IndexAPI interface {
	CreateOne(ctx context.Context, model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
}
