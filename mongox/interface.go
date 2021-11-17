package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IInstance interface {
	Get(filter interface{}, data interface{}) error
	Insert(interface{}) error
	InsertBulk([]*mongo.UpdateOneModel) error
	Delete(filter interface{}) error
	DeleteAll() error
	Update(filter interface{}, update interface{}) error
	Upsert(filter interface{}, update interface{}) error
	FindAndUpdate(filter interface{}, update interface{}, data interface{}) error
	Aggregate(pipeline mongo.Pipeline, data interface{}) error
	Watch() (chan *ChangeEvent, error)
	CreateIndexes(mongo.IndexModel) error
}

type IClient interface {
	Connect() error
	Find(ctx context.Context, database string, collection string, filter interface{}) (ICursor, error)
	InsertOne(ctx context.Context, database string, collection string, value interface{}) (interface{}, error)
	InsertBulk(ctx context.Context, database string, collection string, operations []*mongo.UpdateOneModel) error
	DeleteOne(ctx context.Context, database string, collection string, filter interface{}) (int64, error)
	DeleteMany(ctx context.Context, database string, collection string) (int64, error)
	Update(ctx context.Context, database string, collection string, filter interface{}, update interface{}) (int64, error)
	Upsert(ctx context.Context, database string, collection string, filter interface{}, update interface{}) (interface{}, error)
	FindOneAndUpdate(ctx context.Context, database string, collection string, filter interface{}, update interface{}) (ISingleResult, error)
	Aggregate(ctx context.Context, database string, collection string, pipeline mongo.Pipeline) (ICursor, error)
	Watch(ctx context.Context, database string, collection string) (IChangeStream, error)
	CreateOne(ctx context.Context, database string, collection string, opt mongo.IndexModel) error
}

type ICursor interface {
	Next(ctx context.Context) bool
	All(ctx context.Context, results interface{}) error
	Decode(v interface{}) error
}

type ISingleResult interface {
	Decode(v interface{}) error
}

type IChangeStream interface {
	Next(ctx context.Context) bool
	Current() interface{}
}
