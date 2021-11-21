package kafkax

import (
	"context"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

type ClientConsumerAPI interface {
	SubscribeTopics(topics []string, rebalanceCb cgo.RebalanceCb) error
	Events() chan cgo.Event
	Assign(partitions []cgo.TopicPartition) (err error)
	Unassign() (err error)
}

type ConsumerAPI interface {
	Consume(ctx context.Context) (<-chan Event, error)
}

type ClientProducerAPI interface {
	Close()
	Events() chan cgo.Event
	Produce(msg *cgo.Message, deliveryChan chan cgo.Event) error
	Flush(timeoutMs int) int
}

type ProducerAPI interface {
	Produce(key string, value []byte, opts ...Option) error
}
