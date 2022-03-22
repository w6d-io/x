package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

// CreateIndexes create index based on input mongo index model
func (m *MongoDB) CreateIndexes(opt mongo.IndexModel) error {
	log := logx.WithName(context.TODO(), "Create Index")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	_, err := m.GetCollection().GetIndex().CreateOne(ctx, opt)
	if err != nil {
		log.Error(err, "create index err")
	}
	return err
}

// ListIndexes list indexes
func (m *MongoDB) ListIndexes() ([]string, error) {
	log := logx.WithName(context.TODO(), "List Index")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return nil, errorx.Wrap(err, "fail connect")
	}

	indexList, err := m.GetCollection().GetIndex().ListSpecifications(ctx)
	if err != nil {
		log.Error(err, "list index err")
		return nil, err
	}

	indexes := make([]string, len(indexList))
	for i, v := range indexList {
		indexes[i] = v.Name
	}

	return indexes, nil
}

// DropIndex remove index
func (m *MongoDB) DropIndex(index string) error {
	log := logx.WithName(context.TODO(), "Drop Index")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Connect(); err != nil {
		return errorx.Wrap(err, "fail connect")
	}

	_, err := m.GetCollection().GetIndex().DropOne(ctx, index)
	if err != nil {
		log.Error(err, "drop index err")
	}

	return err
}

// GetIndex return an IndexAPI from collection indexes
func (c *ClientCollection) GetIndex() IndexAPI {
	return &ClientIndex{
		c.Indexes(),
	}
}
