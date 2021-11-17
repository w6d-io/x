package kafkax_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("Controller", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("Client - Mock Producer (json format)", func() {
			ctrl := kafka.MockClient{
				Format: "json",
			}
			err := ctrl.Producer("", `{"test": 1}`)
			Expect(err).To(Succeed())
		})
		It("Controller - Mock Consummer (json format)", func() {
			ctx := context.Background()
			event := kafka.Event{
				Topic: "test_kafka_impl",
				Value: []byte("Hello World!"),
			}
			ctrl := kafka.MockClient{
				Event:  event,
				Format: "json",
			}
			messages, err := ctrl.Consumer(ctx)
			Expect(err).To(Succeed())

			msg := <-messages
			Expect(msg.Topic).To(Equal(event.Topic))
		})
	})
})
