package kafkax

import (
	"context"
	"errors"
	"strings"
	"time"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/w6d-io/x/logx"
)

var (
	ErrConsumer = errors.New("consumer error")
)

func GetConsumerClient(bootstrapServer string, groupId string, username string, password string, opts ...Option) (ClientConsumerAPI, error) {
	options := NewOptions(opts...)

	config := &cgo.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}

	_ = config.SetKey("group.id", groupId)
	_ = config.SetKey("auto.offset.reset", options.AutoOffsetReset)
	_ = config.SetKey("statistics.interval.ms", int(options.StatInterval/time.Millisecond))
	_ = config.SetKey("session.timeout.ms", int(options.SessionTimeout/time.Millisecond))
	_ = config.SetKey("max.poll.interval.ms", int(options.MaxPollInterval/time.Millisecond))
	_ = config.SetKey("go.events.channel.enable", true)
	_ = config.SetKey("go.application.rebalance.enable", options.EnableRebalance)
	_ = config.SetKey("enable.partition.eof", options.EnablePartitionEof)

	if options.GroupInstanceID != "" {
		_ = config.SetKey("group.instance.id", options.GroupInstanceID)
	}
	if len(options.Debugs) > 0 {
		_ = config.SetKey("debug", strings.Join(options.Debugs, ","))
	}
	if options.AuthKafka {
		_ = config.SetKey("sasl.mechanisms", options.Mechanisms)
		_ = config.SetKey("security.protocol", options.Protocol)
		_ = config.SetKey("sasl.username", username)
		_ = config.SetKey("sasl.password", password)
	}

	c, err := cgo.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return &ClientConsumer{
		Consumer: c,
	}, err
}

func (cfg *Kafka) NewConsumer(opts ...Option) (ConsumerAPI, error) {
	clt, err := GetConsumerClient(cfg.BootstrapServer, cfg.GroupId, cfg.Username, cfg.Password, opts...)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		ClientConsumerAPI: clt,
		ListenOnTopics:    cfg.ListenOnTopics,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) (<-chan Event, error) {

	log := logx.WithName(nil, "Consumer")

	err := c.SubscribeTopics(c.ListenOnTopics, nil)
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
				case *cgo.Message:
					event := Event{
						Topic:     string([]byte(*e.TopicPartition.Topic)),
						Key:       e.Key,
						Value:     e.Value,
						Offset:    int64(e.TopicPartition.Offset),
						Partition: e.TopicPartition.Partition,
						Timestamp: e.Timestamp,
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

				case cgo.Error:
					log.Error(ErrConsumer, "Kafka event returns", "code", e.String())

				case cgo.AssignedPartitions:
					log.Info("Assigned Partitions", "code", e.String())
					c.Assign(e.Partitions)

				case cgo.RevokedPartitions:
					log.Info("Revoked Partitions", "code", e.String())
					c.Unassign()

				case cgo.PartitionEOF:
					log.Info("Partition EOF Reached", "code", e.String())

				case cgo.OffsetsCommitted:
					log.Info("OffsetsCommitted", "len", len(e.Offsets))

				default:
					log.V(2).Info("Ignored", "code", e.String())
				}
			}
		}
	}()

	return messages, nil
}
