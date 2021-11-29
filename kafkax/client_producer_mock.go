package kafkax

import (
	"time"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

// MockClientProducer is the internal producer mock client
type MockClientProducer struct {
	ClientProducerAPI
	ErrProduce error
	Event      cgo.Event
}

// Close is an internal mock method
func (p *MockClientProducer) Close() {
}

// Events is an internal mock method
func (p *MockClientProducer) Events() chan cgo.Event {
	evt := make(chan cgo.Event, 1)
	evt <- p.Event
	return evt
}

// Produce is an internal mock method
func (p *MockClientProducer) Produce(msg *cgo.Message, deliveryChan chan cgo.Event) (err error) {
	return p.ErrProduce
}

// Flush is an internal mock method
func (p *MockClientProducer) Flush(timeoutMs int) int {
	time.Sleep(1 * time.Second)
	return 0
}
