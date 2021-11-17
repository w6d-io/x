package kafkax

import (
	"context"
	"errors"
	"strings"
	"time"

	confluent "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/w6d-io/x/logx"
)

var (
	ErrConsumer = errors.New("Consumer Error")
)

func (k *Kafka) Consumer(ctx context.Context, opts ...Option) (<-chan Event, error) {

	log := logx.WithName(nil, "Consumer")
	options := NewOptions(opts...)

	config := &confluent.ConfigMap{
		"bootstrap.servers": k.BootstrapServer,
	}

	config.SetKey("group.id", k.GroupId)
	config.SetKey("auto.offset.reset", options.AutoOffsetReset)
	config.SetKey("statistics.interval.ms", int(options.StatInterval/time.Millisecond))
	config.SetKey("session.timeout.ms", int(options.SessionTimeout/time.Millisecond))
	config.SetKey("max.poll.interval.ms", int(options.MaxPollInterval/time.Millisecond))
	config.SetKey("go.events.channel.enable", true)
	config.SetKey("go.application.rebalance.enable", options.EnableRebalance)
	config.SetKey("enable.partition.eof", options.EnablePartitionEof)

	if options.GroupInstanceID != "" {
		config.SetKey("group.instance.id", options.GroupInstanceID)
	}
	if len(options.Debugs) > 0 {
		config.SetKey("debug", strings.Join(options.Debugs, ","))
	}
	if options.AuthKafka {
		config.SetKey("sasl.mechanisms", options.Mechanisms)
		config.SetKey("security.protocol", options.Protocol)
		config.SetKey("sasl.username", k.Username)
		config.SetKey("sasl.password", k.Password)
	}

	c, err := confluent.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics(k.ListenOnTopics, nil)
	if err != nil {
		return nil, err
	}

	messages := make(chan Event)

	go func() {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				log.Info("Ctx.Done()")
				return
			case ev := <-c.Events():
				switch e := ev.(type) {
				case *confluent.Message:
					event := Event{
						Topic:     string([]byte(*e.TopicPartition.Topic)),
						Key:       e.Key,
						Value:     e.Value,
						Offset:    int64(e.TopicPartition.Offset),
						Partition: e.TopicPartition.Partition,
						Timestamp: e.Timestamp,
					}
					if k.SchemaRegistry != nil {
						event.SchemaRegistry = k.SchemaRegistry
					}
					if e.Headers != nil {
						for _, h := range e.Headers {
							header := Header{
								Key:   h.Key,
								Value: h.Value,
							}
							event.Headers = append(event.Headers, header)
						}
					}
					messages <- event

				case confluent.Error:
					log.Error(ErrConsumer, "Kafka event returns", "code", e.String())

				case confluent.AssignedPartitions:
					log.Info("Assigned Partitions", "code", e.String())
					c.Assign(e.Partitions)

				case confluent.RevokedPartitions:
					log.Info("Revoked Partitions", "code", e.String())
					c.Unassign()

				case confluent.PartitionEOF:
					log.Info("Partition EOF Reached", "code", e.String())

				case confluent.OffsetsCommitted:
					log.Info("OffsetsCommitted", "len", len(e.Offsets))

				default:
					log.Info("Ignored", "code", e.String())
				}
			}
		}
	}()

	return messages, nil
}
