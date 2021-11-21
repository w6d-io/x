package kafkax

import (
	"errors"
	"time"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/w6d-io/x/logx"
)

var (
	ErrProducer = errors.New("Producer Error")
)

func GetProducerClient(bootstrapServer string, username string, password string, opts ...Option) (ClientProducerAPI, error) {
	options := NewOptions(opts...)

	config := &cgo.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}
	if options.AuthKafka {
		_ = config.SetKey("sasl.mechanisms", options.Mechanisms)
		_ = config.SetKey("security.protocol", options.Protocol)
		_ = config.SetKey("sasl.username", username)
		_ = config.SetKey("sasl.password", password)
	}

	p, err := cgo.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return &ClientProducer{
		Producer: p,
	}, err
}

func (cfg *Kafka) NewProducer(opts ...Option) (ProducerAPI, error) {
	clt, err := GetProducerClient(cfg.BootstrapServer, cfg.Username, cfg.Password, opts...)
	if err != nil {
		return nil, err
	}
	return &Producer{
		ClientProducerAPI: clt,
		ProducToTopic:     cfg.ProducToTopic,
	}, nil
}

func (p *Producer) Produce(key string, value []byte, opts ...Option) error {

	log := logx.WithName(nil, "Producer")

	options := NewOptions(opts...)

	go func() {
		defer p.Close()
		for e := range p.Events() {
			switch ev := e.(type) {
			case *cgo.Message:
				if ev.TopicPartition.Error != nil {
					log.Error(ev.TopicPartition.Error, "Failed to deliver",
						"stacktrace", ev.TopicPartition)
				} else {
					log.Info("Successfully produced record",
						"topic", *ev.TopicPartition.Topic,
						"partition", ev.TopicPartition.Partition,
						"offset", ev.TopicPartition.Offset)
				}
			}
		}
	}()

	if err := p.ClientProducerAPI.Produce(&cgo.Message{
		TopicPartition: cgo.TopicPartition{Topic: &p.ProducToTopic, Partition: cgo.PartitionAny},
		Key:            []byte(key),
		Value:          value,
		Timestamp:      time.Now(),
	}, nil); err != nil {
		log.Error(err, "produce failed")
		return err
	}

	p.Flush(int(options.WriteTimeout / time.Millisecond))

	return nil

}
