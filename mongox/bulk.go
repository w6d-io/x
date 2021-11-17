package mongox

import "go.mongodb.org/mongo-driver/mongo"

func BuildBulkOperation(filter interface{}, update interface{}, data interface{}) *mongo.UpdateOneModel {

	operation := mongo.NewUpdateOneModel()
	operation.SetFilter(filter)
	operation.SetUpdate(update)
	operation.SetUpsert(true)

	return operation
}
