package kafkax

import (
	"context"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

// ClientConsumerAPI is the internal client consumer API
type ClientConsumerAPI interface {
	Unsubscribe() error
	SubscribeTopics([]string, cgo.RebalanceCb) error
	Events() chan cgo.Event
	Assign([]cgo.TopicPartition) error
	Unassign() error
}

// ConsumerAPI is the public client consumer API
type ConsumerAPI interface {
	SetTopics(...string) (ConsumerAPI, error)
	GetTopics() []string
	Consume(context.Context) (<-chan Event, error)
}

// ClientProducerAPI is the internal client producer API
type ClientProducerAPI interface {
	Close()
	Events() chan cgo.Event
	Produce(*cgo.Message, chan cgo.Event) error
	Flush(int) int
}

// ProducerAPI is the public client producer API
type ProducerAPI interface {
	SetTopic(string) ProducerAPI
	GetTopic() string
	Produce(string, []byte, ...Option) error
}
