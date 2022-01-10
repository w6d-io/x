package mongox

import (
	mgo "go.mongodb.org/mongo-driver/mongo"
)

// Mongo input structure
type Mongo struct {
	// Username for database authentication
	Username string `mapstructure:"username"`
	// Password for database authentication
	Password string `mapstructure:"password"`
	// Name of the database
	Name string `mapstructure:"name"`
	// URL of the database
	URL string `mapstructure:"url"`
	// Name of the authorisation database
	AuthSource string `mapstructure:"authSource"`
}

// Client is the internal Mongo Client
type Client struct {
	*mgo.Client
	Database   string
	Collection string
}

// ClientDatabase is the internal Mongo Database
type ClientDatabase struct {
	*mgo.Database
}

// ClientCollection is the internal Mongo Collection
type ClientCollection struct {
	*mgo.Collection
}

// ClientCursor is the internal Mongo Cursor
type ClientCursor struct {
	*mgo.Cursor
}

// ClientSingleResult is the internal Single Result
type ClientSingleResult struct {
	*mgo.SingleResult
}

// ClientIndex is the internal Client Index
type ClientIndex struct {
	mgo.IndexView
}

// MongoDB is the public Mongo Instance
type MongoDB struct {
	ClientAPI
	// Keep internal connection status
	isConnected bool
	// From configuration section
	cfg *Mongo
	// Options
	options *Options
	// From user selection
	Collection string
}
