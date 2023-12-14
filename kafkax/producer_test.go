//go:build !integration

package kafkax_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cgo "github.com/confluentinc/confluent-kafka-go/kafka"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("Producer", func() {
	Context("options", func() {
		It("bad protocol configuration", func() {
			var opts []kafka.Option
			opts = append(opts, kafka.Protocol("Unknown"))
			opts = append(opts, kafka.AuthKafka(true))
			_ = kafka.NewOptions(opts...)
			k := kafka.Kafka{
				Username:        "test",
				Password:        "test",
				BootstrapServer: "localhost:9092",
			}
			_, err := k.NewProducer(opts...)
			Expect(err).ToNot(Succeed())
			Expect(err.Error()).To(Equal("Invalid value \"Unknown\" for configuration property \"security.protocol\""))
		})
		It("", func() {
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
			_, err := k.NewProducer(opts...)
			Expect(err).To(Succeed())
		})
	})
	Context("produce", func() {
		It("success", func() {
			client := &kafka.MockClientProducer{}
			pm := kafka.Producer{
				ClientProducerAPI: client,
			}
			header := []kafka.Header{{
				Key:   "key",
				Value: []byte("hello world"),
			}}
			err := pm.SetTopic("test").Produce("key", []byte(string("Hello World")), kafka.WithHeaders(header))
			Expect(err).To(Succeed())
		})
		It("missing topic while producing", func() {
			client := &kafka.MockClientProducer{}
			pm := kafka.Producer{
				ClientProducerAPI: client,
			}
			err := pm.Produce("key", []byte(string("Hello World")))
			Expect(err).NotTo(Succeed())
		})
		It("fails while producing", func() {
			client := &kafka.MockClientProducer{
				ErrProduce: errors.New("fail while producing"),
			}
			pm := kafka.Producer{
				ClientProducerAPI: client,
			}
			err := pm.SetTopic("test").Produce("key", []byte(string("Hello World")))
			Expect(err).NotTo(Succeed())
		})
		It("partition error while producing", func() {
			client := &kafka.MockClientProducer{
				Event: &cgo.Message{
					TopicPartition: cgo.TopicPartition{
						Error: errors.New("partition error"),
					},
				},
			}
			pm := kafka.Producer{
				ClientProducerAPI: client,
			}
			err := pm.SetTopic("test").Produce("key", []byte(string("Hello World")))
			Expect(err).To(Succeed())
		})
		It("success while producing", func() {
			var topic string = "test"
			var partition int32 = 1
			var offset cgo.Offset = cgo.Offset(1)
			client := &kafka.MockClientProducer{
				Event: &cgo.Message{
					TopicPartition: cgo.TopicPartition{
						Topic:     &topic,
						Partition: partition,
						Offset:    offset,
					},
				},
			}
			pm := kafka.Producer{
				ClientProducerAPI: client,
			}
			err := pm.SetTopic("test").Produce("key", []byte(string("Hello World")))
			Expect(err).To(Succeed())
			pm.Close()
		})
	})
})
