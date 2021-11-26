package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/w6d-io/x/errorx"
	"github.com/w6d-io/x/logx"
)

func (m *MongoDB) Incr(key string) (int64, error) {
	log := logx.WithName(nil, "Incr")
	if err := m.Connect(); err != nil {
		return -1, errorx.Wrap(err, "fail connect")
	}
	type p struct {
		Id int64 `bson:"id"`
	}
	var pid []p
	cursor, err := m.Get(bson.M{"_id": key})
	if err != nil && err != mongo.ErrNoDocuments {
		return 0, err
	}
	if err := cursor.All(context.Background(), &pid); err != nil {
		return -1, errorx.Wrap(err, "fail connect")
	}
	if len(pid) == 0 {
		pid = append(pid, p{Id: 1})
		err = m.Insert(bson.M{"_id": key, "id": pid[0].Id})
		if err != nil {
			log.Error(err, "insert")
			return 0, err
		}
	}
	err = m.Update(bson.M{"_id": key}, bson.M{"$inc": bson.M{"id": 1}})
	if err != nil {
		return 0, err
	}
	log.WithValues("data", pid[0].Id).V(1).Info("result from incr")
	return pid[0].Id, nil
}
