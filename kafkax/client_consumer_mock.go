package kafkax

import (
	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

type MockClientConsumer struct {
	ClientConsumerAPI
	ErrSubscribeTopics   error
	ErrAssignPartition   error
	ErrUnAssignPartition error
	Event                cgo.Event
}

func (c *MockClientConsumer) SubscribeTopics(topics []string, rebalanceCb cgo.RebalanceCb) (err error) {
	return c.ErrSubscribeTopics
}

func (c *MockClientConsumer) Events() chan cgo.Event {
	evt := make(chan cgo.Event, 1)
	evt <- c.Event
	return evt
}

func (c *MockClientConsumer) Assign(partitions []cgo.TopicPartition) (err error) {
	return c.ErrAssignPartition
}

func (c *MockClientConsumer) Unassign() (err error) {
	return c.ErrUnAssignPartition
}
