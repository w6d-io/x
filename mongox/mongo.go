package mongox

import (
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	// Username for database authentication
	Username string `mapstructure:"username"`
	// Password for database authentication
	Password string `mapstructure:"password"`
	// Name of the database
	Name string `mapstructure:"name"`
	// URL of the database
	URL string `mapstructure:"url"`
	// URL of the database
	AuthSource string `mapstructure:"authSource"`
}

type Client struct {
	*mgo.Client
	Database   string
	Collection string
	opts       []*options.ClientOptions
}

type ClientDatabase struct {
	*mgo.Database
}

type ClientCollection struct {
	*mgo.Collection
}

type ClientCursor struct {
	*mgo.Cursor
}

type ClientSingleResult struct {
	*mgo.SingleResult
}

type ClientIndex struct {
	mgo.IndexView
}

type MongoDB struct {
	ClientAPI
	isConnected bool
	Database    string
	Collection  string
}
