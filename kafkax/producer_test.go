package kafkax_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("Kafka", func() {
	Context("options", func() {
		It("bad configuration", func() {
			var opts []kafka.Option
			opts = append(opts, kafka.Protocol("SASL_SSL"))
			opts = append(opts, kafka.Mechanisms("PLAIN"))
			opts = append(opts, kafka.Async(false))
			opts = append(opts, kafka.WriteTimeout(1*time.Second))
			opts = append(opts, kafka.MaxWait(1*time.Second))
			opts = append(opts, kafka.StatInterval(3*time.Second))
			opts = append(opts, kafka.NumPartitions(10))
			opts = append(opts, kafka.ReplicationFactor(3))
			opts = append(opts, kafka.AuthKafka(true))
			opts = append(opts, kafka.FullStats(false))
			opts = append(opts, kafka.Debugs([]string{"deep"}))
			opts = append(opts, kafka.SessionTimeout(10*time.Millisecond))
			opts = append(opts, kafka.MaxPollInterval(10*time.Millisecond))
			opts = append(opts, kafka.GroupInstanceID("test"))
			opts = append(opts, kafka.ConfigMapKey("test"))
			_ = kafka.NewOptions(opts...)
			k := kafka.Kafka{
				Username:        "test",
				Password:        "test",
				BootstrapServer: "localhost:9092",
			}
			err := k.Producer("TEST_ID", "message", opts...)
			Expect(err).ToNot(Succeed())
			Expect(err.Error()).To(Equal("Local: Invalid argument or configuration"))
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
				ProducToTopic:   "TEST",
				Username:        "test",
				Password:        "test",
				BootstrapServer: "localhost:9092",
			}
			err := k.Producer("TEST_ID", "message", opts...)
			Expect(err).To(Succeed())
		})
	})
})
