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
	// ErrConsumer ...
	ErrConsumer = errors.New("consumer error")
	// ErrConsumerTopicsRequest ...
	ErrConsumerTopicsRequest = errors.New("topics request is not registered")
	// ErrConsumerTopicsIsNotSet ...
	ErrConsumerTopicsIsNotSet = errors.New("topics for consumer message is not set")
)

// GetConsumerClient returns a ClientConsumerAPI
func GetConsumerClient(bootstrapServer string, groupID string, username string, password string, opts ...Option) (ClientConsumerAPI, error) {
	options := NewOptions(opts...)

	config := &cgo.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}

	_ = config.SetKey("group.id", groupID)
	_ = config.SetKey("auto.offset.reset", options.AutoOffsetReset)
	_ = config.SetKey("statistics.interval.ms", int(options.StatInterval/time.Millisecond))
	_ = config.SetKey("session.timeout.ms", int(options.SessionTimeout/time.Millisecond))
	_ = config.SetKey("max.poll.interval.ms", int(options.MaxPollInterval/time.Millisecond))
	_ = config.SetKey("go.events.channel.enable", true)
	_ = config.SetKey("go.application.rebalance.enable", options.EnableRebalance)
	_ = config.SetKey("enable.partition.eof", options.EnablePartitionEOF)

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

// NewConsumer creates a ConsumerAPI
func (cfg *Kafka) NewConsumer(groupID string, opts ...Option) (ConsumerAPI, error) {
	clt, err := GetConsumerClient(cfg.BootstrapServer, groupID, cfg.Username, cfg.Password, opts...)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		ClientConsumerAPI: clt,
	}, nil
}

// SetTopics assign topics to the consumer
func (c *Consumer) SetTopics(topics ...string) (ConsumerAPI, error) {
	log := logx.WithName(context.TODO(), "Set topics")

	if c.topicsReqChan == nil {
		c.topicsReqChan = make(chan []string, 1)
	}

	select {
	case c.topicsReqChan <- topics:
	default:
		log.Error(ErrConsumerTopicsRequest, "there is already pending request")
		return nil, ErrConsumerTopicsRequest
	}

	return c, nil
}

// GetTopics return topics assigned to the consumer
func (c *Consumer) GetTopics() []string {
	return c.topics
}

// Consume consumer message for consumer
//gocyclo:ignore
func (c *Consumer) Consume(ctx context.Context) (<-chan Event, error) {

	log := logx.WithName(context.TODO(), "Consumer")

	if c.topicsReqChan == nil {
		return nil, ErrConsumerTopicsIsNotSet
	}

	messages := make(chan Event)

	go func() {
		defer close(messages)
		for {
			select {
			case <-ctx.Done():
				log.Info("Ctx.Done()")
				return
			case topics := <-c.topicsReqChan:
				err := c.Unsubscribe()
				if err != nil {
					log.Error(err, "error while unsubscribing")
					return
				}
				err = c.SubscribeTopics(topics, nil)
				if err != nil {
					log.Error(err, "error while subscribing")
					return
				}
				c.topics = topics
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
					_ = c.Assign(e.Partitions)

				case cgo.RevokedPartitions:
					log.Info("Revoked Partitions", "code", e.String())
					_ = c.Unassign()

				case cgo.PartitionEOF:
					log.Info("Partition EOF Reached", "code", e.String())

				case cgo.OffsetsCommitted:
					log.Info("OffsetsCommitted", "len", len(e.Offsets))

				default:
					log.V(3).Info("Ignored", "code", e)
				}
			}
		}
	}()

	return messages, nil
}
