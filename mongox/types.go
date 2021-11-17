package mongox

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Instance struct {
	Client     IClient
	Database   string
	Collection string
}

type Client struct {
	// mongo client
	client *mongo.Client
	// local mongo options storage
	opts []*options.ClientOptions
}

type Cursor struct {
	cursor *mongo.Cursor
}

type SingleResult struct {
	singleResult *mongo.SingleResult
}

type ChangeStream struct {
	changeStream *mongo.ChangeStream
}

var (
	db *Client
)
