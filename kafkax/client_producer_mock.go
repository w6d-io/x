package kafkax

import (
	"time"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

type MockClientProducer struct {
	ClientProducerAPI
	ErrProduce error
	Event      cgo.Event
}

func (p *MockClientProducer) Close() {
}

func (p *MockClientProducer) Events() chan cgo.Event {
	evt := make(chan cgo.Event, 1)
	evt <- p.Event
	return evt
}

func (p *MockClientProducer) Produce(msg *cgo.Message, deliveryChan chan cgo.Event) (err error) {
	return p.ErrProduce
}

func (p *MockClientProducer) Flush(timeoutMs int) int {
	time.Sleep(1 * time.Second)
	return 0
}
