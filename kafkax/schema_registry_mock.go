package kafkax

type MockSchemaRegistry struct {
	Schema Schema
}

func (sr *MockSchemaRegistry) GetLatestSchema(topic string) (*Schema, error) {
	return &sr.Schema, nil
}

var (
	_ ISchemaRegistry = &MockSchemaRegistry{}
)
