package kafkax

import (
	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

// MockClientConsumer is the internal consumer mock client
type MockClientConsumer struct {
	ClientConsumerAPI
	ErrUnsubscribe       error
	ErrSubscribeTopics   error
	ErrAssignPartition   error
	ErrUnAssignPartition error
	Event                cgo.Event
}

// Unsubscribe is an internal mock method
func (c *MockClientConsumer) Unsubscribe() (err error) {
	return c.ErrUnsubscribe
}

// SubscribeTopics is an internal mock method
func (c *MockClientConsumer) SubscribeTopics(topics []string, rebalanceCb cgo.RebalanceCb) (err error) {
	return c.ErrSubscribeTopics
}

// Events is an internal mock method
func (c *MockClientConsumer) Events() chan cgo.Event {
	evt := make(chan cgo.Event, 1)
	evt <- c.Event
	return evt
}

// Assign is an internal mock method
func (c *MockClientConsumer) Assign(partitions []cgo.TopicPartition) (err error) {
	return c.ErrAssignPartition
}

// Unassign is an internal mock method
func (c *MockClientConsumer) Unassign() (err error) {
	return c.ErrUnAssignPartition
}
