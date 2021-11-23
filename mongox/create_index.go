package mongox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (m *MongoDB) CreateIndexes(opt mongo.IndexModel) error {
	log := logx.WithName(nil, "Create Index")
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

func (c *ClientCollection) GetIndex() IndexAPI {
	return &ClientIndex{
		c.Indexes(),
	}
}
