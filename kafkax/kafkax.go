package kafkax

import (
	cgo "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	BootstrapServer string   `json:"boostrapserver" mapstructure:"boostrapserver"`
	Username        string   `json:"username" mapstructure:"username"`
	Password        string   `json:"password" mapstructure:"password"`
	GroupId         string   `json:"groupid" mapstructure:"groupid"`
	ListenOnTopics  []string `json:"listenontopics" mapstructure:"listenontopics"`
	ProducToTopic   string   `json:"productotopic" mapstructure:"productotopic"`
}

type ClientConsumer struct {
	*cgo.Consumer
}

type ClientProducer struct {
	*cgo.Producer
}

type Consumer struct {
	ClientConsumerAPI
	ListenOnTopics []string
}

type Producer struct {
	ClientProducerAPI
	ProducToTopic string
}
