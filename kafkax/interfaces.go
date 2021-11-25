package kafkax

import (
	"context"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

type ClientConsumerAPI interface {
	Unsubscribe() error
	SubscribeTopics([]string, cgo.RebalanceCb) error
	Events() chan cgo.Event
	Assign([]cgo.TopicPartition) error
	Unassign() error
}

type ConsumerAPI interface {
	SetTopics(...string) (ConsumerAPI, error)
	GetTopics() []string
	Consume(context.Context) (<-chan Event, error)
}

type ClientProducerAPI interface {
	Close()
	Events() chan cgo.Event
	Produce(*cgo.Message, chan cgo.Event) error
	Flush(int) int
}

type ProducerAPI interface {
	SetTopic(string) ProducerAPI
	GetTopic() string
	Produce(string, []byte, ...Option) error
}
