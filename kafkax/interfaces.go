package kafkax

import (
	"context"
)

type IClient interface {
	Consumer(ctx context.Context, opts ...Option) (<-chan Event, error)
	Producer(key string, value interface{}, opts ...Option) error
}

type ISchemaRegistry interface {
	GetLatestSchema(topic string) (*Schema, error)
}

var (
	_ IClient         = &Kafka{}
	_ ISchemaRegistry = &SchemaRegistry{}
)
