package kafkax

import (
	"errors"
	"time"

	confluent "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/w6d-io/x/logx"
)

var (
	ErrProducer = errors.New("Producer Error")
)

func (k *Kafka) Producer(key string, value interface{}, opts ...Option) error {

	log := logx.WithName(nil, "Producer")

	options := NewOptions(opts...)

	config := &confluent.ConfigMap{
		"bootstrap.servers": k.BootstrapServer,
	}
	if options.AuthKafka {
		_ = config.SetKey("sasl.mechanisms", options.Mechanisms)
		_ = config.SetKey("security.protocol", options.Protocol)
		_ = config.SetKey("sasl.username", k.Username)
		_ = config.SetKey("sasl.password", k.Password)
	}

	p, err := confluent.NewProducer(config)
	if err != nil {
		return err
	}

	go func() {
		defer p.Close()
		for e := range p.Events() {
			switch ev := e.(type) {
			case *confluent.Message:
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

	event := Event{
		Topic:          k.ProducToTopic,
		SchemaRegistry: k.SchemaRegistry,
	}
	bValue, err := event.Marshall(value)
	if err != nil {
		return err
	}

	if err = p.Produce(&confluent.Message{
		TopicPartition: confluent.TopicPartition{Topic: &k.ProducToTopic, Partition: confluent.PartitionAny},
		Key:            []byte(key),
		Value:          bValue,
		Timestamp:      time.Now(),
	}, nil); err != nil {
		log.Error(err, "produce failed")
		return err
	}

	p.Flush(int(options.WriteTimeout / time.Millisecond))

	return nil

}
