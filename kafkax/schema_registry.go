package kafkax

import (
	"fmt"

	"github.com/riferrei/srclient"
)

func (s *SchemaRegistry) GetLatestSchema(topic string) (*Schema, error) {
	if s.Url == "" {
		return nil, nil
	}
	client := srclient.CreateSchemaRegistryClient(s.Url)
	pattern := "%s"
	if s.TopicPattern != "" {
		pattern = s.TopicPattern
	}
	valueSchema, err := client.GetLatestSchema(fmt.Sprintf(pattern, topic))
	if err != nil {
		return &Schema{}, err
	}

	return &Schema{
		Id:     valueSchema.ID(),
		Schema: valueSchema.Schema(),
	}, nil
}
