package kafkax

import (
	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	BootstrapServer string `json:"boostrapserver" mapstructure:"boostrapserver"`
	Username        string `json:"username" mapstructure:"username"`
	Password        string `json:"password" mapstructure:"password"`
	GroupId         string `json:"groupid" mapstructure:"groupid"`
}

type ClientConsumer struct {
	*cgo.Consumer
}

type ClientProducer struct {
	*cgo.Producer
}

type Consumer struct {
	ClientConsumerAPI
	topicsReqChan chan []string
	topics        []string
}

type Producer struct {
	ClientProducerAPI
	topic string
}
