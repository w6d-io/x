package mongox

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}

// ChangeEvent document according
// https://docs.mongodb.com/manual/reference/change-events/#change-stream-output
type ChangeEvent struct {
	ID                interface{} `bson:"_id"`
	Operation         string      `bson:"operationType"`
	Document          bson.M      `bson:"fullDocument"`
	Namespace         bson.M      `bson:"ns"`
	NewCollectionName bson.M      `bson:"to,omitempty"`
	DocumentKey       documentKey `bson:"documentKey"`
	Updates           bson.M      `bson:"updateDescription,omitempty"`
	ClusterTime       time.Time   `bson:"clusterTime"`
	Transaction       int64       `bson:"txnNumber,omitempty"`
	SessionID         bson.M      `bson:"lsid,omitempty"`
}
