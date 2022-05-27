package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoAPI is the public Mongo API interface
type MongoAPI interface {
	SetCollection(string) MongoAPI
	SetOptions(...Option) MongoAPI
	Connect() error
	Get(interface{}, interface{}, ...*options.FindOptions) error
	Insert(interface{}) error
	InsertBulk([]*mongo.UpdateOneModel) error
	Delete(interface{}) error
	DeleteAll() error
	Update(interface{}, interface{}) error
	Upsert(interface{}, interface{}) error
	FindAndUpdate(interface{}, interface{}, interface{}) error
	Aggregate(mongo.Pipeline, interface{}) error
	CreateIndexes(mongo.IndexModel) error
	CreateManyIndexes([]mongo.IndexModel) error
	ListIndexes() ([]string, error)
	DropIndex(string) error
	Incr(string) (int64, error)
}

// ClientAPI is the internal Client API interface
type ClientAPI interface {
	SetCollection(collection string)
	GetCollection() CollectionAPI
	SetCursor(*mongo.Cursor) CursorAPI
	SetSingleResult(*mongo.SingleResult) SingleResultAPI
	Connect(context.Context) error
	Ping(context.Context, *readpref.ReadPref) error
}

// CollectionAPI is the internal Collection API interface
type CollectionAPI interface {
	GetIndex() IndexAPI
	InsertOne(context.Context, interface{}, ...*mgoOtions.InsertOneOptions) (*mongo.InsertOneResult, error)
	BulkWrite(context.Context, []mongo.WriteModel, ...*mgoOtions.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	Find(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(context.Context, interface{}, ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	ReplaceOne(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	FindOneAndUpdate(context.Context, interface{}, interface{}, ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	Aggregate(context.Context, interface{}, ...*options.AggregateOptions) (*mongo.Cursor, error)
}

// CursorAPI is the internal Cursor API interface
type CursorAPI interface {
	Next(context.Context) bool
	All(context.Context, interface{}) error
	Decode(interface{}) error
}

// SingleResultAPI is the internal Single Result API interface
type SingleResultAPI interface {
	Decode(interface{}) error
}

// IndexAPI is the internal Index API interface
type IndexAPI interface {
	ListSpecifications(ctx context.Context, opts ...*options.ListIndexesOptions) ([]*mongo.IndexSpecification, error)
	DropOne(ctx context.Context, name string, opts ...*options.DropIndexesOptions) (bson.Raw, error)
	CreateOne(context.Context, mongo.IndexModel, ...*options.CreateIndexesOptions) (string, error)
	CreateMany(context.Context, []mongo.IndexModel, ...*options.CreateIndexesOptions) ([]string, error)
}
