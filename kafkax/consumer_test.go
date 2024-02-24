//go:build !integration

package kafkax_test

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("Consumer", func() {
	Context("options", func() {
		It("bad protocol configuration", func() {
			var opts []kafka.Option
			opts = append(opts, kafka.Protocol("Unknown"))
			opts = append(opts, kafka.AuthKafka(true))
			opts = append(opts, kafka.GroupInstanceID("test"))
			opts = append(opts, kafka.MaxBytes(1))
			opts = append(opts, kafka.MinBytes(1))
			opts = append(opts, kafka.NumPartitions(1))
			opts = append(opts, kafka.ReplicationFactor(1))
			opts = append(opts, kafka.FullStats(false))
			opts = append(opts, kafka.ConfigMapKey("test"))
			opts = append(opts, kafka.Rebalance(false))
			opts = append(opts, kafka.PartitionEOF(false))
			opts = append(opts, kafka.EarliestOffset())
			_ = kafka.NewOptions(opts...)
			k := kafka.Kafka{
				Username:        "test",
				Password:        "test",
				BootstrapServer: "localhost:9092",
			}
			_, err := k.NewConsumer("groupID", opts...)
			Expect(err).ToNot(Succeed())
			Expect(err.Error()).To(Equal("Invalid value \"Unknown\" for configuration property \"security.protocol\""))
		})
		It("bad debug configuration", func() {
			var opts []kafka.Option
			opts = append(opts, kafka.Debugs([]string{"deep"}))
			_ = kafka.NewOptions(opts...)
			k := kafka.Kafka{
				Username:        "test",
				Password:        "test",
				BootstrapServer: "localhost:9092",
			}
			_, err := k.NewConsumer("groupID", opts...)
			Expect(err).ToNot(Succeed())
			Expect(err.Error()).To(Equal("Invalid value \"deep\" for configuration property \"debug\""))
		})
		It("success to configure a new consumer", func() {
			var opts []kafka.Option
			opts = append(opts, kafka.Protocol("SASL_SSL"))
			opts = append(opts, kafka.Mechanisms("PLAIN"))
			opts = append(opts, kafka.Async(false))
			opts = append(opts, kafka.WriteTimeout(1*time.Second))
			opts = append(opts, kafka.MaxWait(1*time.Second))
			opts = append(opts, kafka.StatInterval(3*time.Second))
			opts = append(opts, kafka.AuthKafka(true))
			opts = append(opts, kafka.SessionTimeout(10*time.Millisecond))
			opts = append(opts, kafka.MaxPollInterval(10*time.Millisecond))
			_ = kafka.NewOptions(opts...)
			k := kafka.Kafka{
				Username:        "test",
				Password:        "test",
				BootstrapServer: "localhost:9092",
			}
			_, err := k.NewConsumer("groupID", opts...)
			Expect(err).To(Succeed())
		})
	})
	Context("consume", func() {
		It("success", func() {
			client := &kafka.MockClientConsumer{
				Event: &cgo.Stats{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("set topic and read back", func() {
			client := &kafka.MockClientConsumer{
				Event: &cgo.Stats{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
			Expect(cm.GetTopics()).To(Equal([]string{"test"}))
		})
		It("multiple set topics and read back", func() {
			client := &kafka.MockClientConsumer{
				Event: &cgo.Stats{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.SetTopics("test")
			Expect(err).NotTo(Succeed())
		})
		It("missing topic registration", func() {
			client := &kafka.MockClientConsumer{}
			cm := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			_, err := cm.Consume(context.Background())
			Expect(err).NotTo(Succeed())
		})
		It("fails while subscribing", func() {
			client := &kafka.MockClientConsumer{
				ErrSubscribeTopics: errors.New("fail while subscribing"),
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(context.Background())
			Expect(err).To(Succeed())
		})
		It("read event while consuming", func() {
			var topic string = "test"
			var partition int32 = 1
			var offset cgo.Offset = cgo.Offset(1)
			client := &kafka.MockClientConsumer{
				Event: &cgo.Message{
					TopicPartition: cgo.TopicPartition{
						Topic:     &topic,
						Partition: partition,
						Offset:    offset,
					},
					Value:     []byte("hello world"),
					Key:       []byte("key"),
					Timestamp: time.Now(),
					Headers: []cgo.Header{{
						Value: []byte("hello world"),
						Key:   "key",
					}},
				},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("error event while consuming", func() {
			client := &kafka.MockClientConsumer{
				Event: cgo.Error{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("assigned partition event while consuming", func() {
			client := &kafka.MockClientConsumer{
				Event: cgo.AssignedPartitions{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("revoked partition event while consuming", func() {
			client := &kafka.MockClientConsumer{
				Event: cgo.RevokedPartitions{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("eof partition event while consuming", func() {
			client := &kafka.MockClientConsumer{
				Event: cgo.PartitionEOF{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("committed partition event while consuming", func() {
			client := &kafka.MockClientConsumer{
				Event: cgo.OffsetsCommitted{},
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
		It("unsubscribe error while consuming", func() {
			client := &kafka.MockClientConsumer{
				Event:          cgo.OffsetsCommitted{},
				ErrUnsubscribe: errors.New("unsubscribe error"),
			}
			clientCons := kafka.Consumer{
				ClientConsumerAPI: client,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cm, err := clientCons.SetTopics("test")
			Expect(err).To(Succeed())
			_, err = cm.Consume(ctx)
			time.Sleep(1 * time.Second)
			cancel()
			Expect(err).To(Succeed())
		})
	})
})
