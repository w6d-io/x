package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (i *Instance) Watch() (chan *ChangeEvent, error) {

	log := logx.WithName(nil, "Watch")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := i.Client.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}

	// opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	collectionStream, err := i.Client.Watch(ctx, i.Database, i.Collection)
	if err != nil {
		log.Error(err, "watch err")
		return nil, err
	}

	var events = make(chan *ChangeEvent)
	go func() {
		defer close(events)
		for collectionStream.Next(ctx) {
			var ce ChangeEvent
			err := bson.Unmarshal(collectionStream.Current().([]byte), &ce)
			if err == nil {
				events <- &ce
			}
		}
	}()

	return events, nil
}

func (c *Client) Watch(ctx context.Context, database string, collection string) (IChangeStream, error) {
	col := c.client.Database(database).Collection(collection)

	changeStream, err := col.Watch(ctx, mongo.Pipeline{})

	if err != nil {
		return nil, err
	}
	return &ChangeStream{
		changeStream: changeStream,
	}, nil
}
