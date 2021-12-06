package kafkax

import (
	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Kafka input structure
type Kafka struct {
	BootstrapServer string `json:"boostrapserver" mapstructure:"boostrapserver"`
	Username        string `json:"username" mapstructure:"username"`
	Password        string `json:"password" mapstructure:"password"`
}

// ClientConsumer is the internal Kafka client consumer
type ClientConsumer struct {
	*cgo.Consumer
}

// ClientProducer is the internal Kafka client producer
type ClientProducer struct {
	*cgo.Producer
}

// Consumer is the public Kafka client consumer
type Consumer struct {
	ClientConsumerAPI
	topicsReqChan chan []string
	topics        []string
}

// Producer is the public Kafka client producer
type Producer struct {
	ClientProducerAPI
	topic string
}
