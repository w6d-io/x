package mongox

import (
	"bytes"
	"text/template"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePipelineFromTemplate(
	aggTemplate string,
	fields interface{},
) (mongo.Pipeline, error) {

	t := template.Must(template.New("agg").Parse(aggTemplate))

	// Execute the template.
	var tpl bytes.Buffer
	err := t.Execute(&tpl, fields)
	if err != nil {
		return nil, err
	}

	var bsonMap interface{}
	err = bson.UnmarshalExtJSON(tpl.Bytes(), false, &bsonMap)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{}

	for _, bsonA := range bsonMap.(bson.A) {
		pipeline = append(pipeline, bsonA.(bson.D))
	}

	return pipeline, nil
}
