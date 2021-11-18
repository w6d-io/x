package kafkax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("Event Marshaller", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("Publish (json format)", func() {
			event := kafka.Event{}
			msg := TTest{
				Message: "Hello World!",
			}
			_, err := event.Marshall(msg)
			Expect(err).To(Succeed())
		})
		It("Publish a malformated (json format)", func() {
			event := kafka.Event{}
			_, err := event.Marshall(make(chan int))
			Expect(err).NotTo(Succeed())
		})
		It("Publish (proto format)", func() {
			event := kafka.Event{}
			msg := Test{
				Message: "Hello World!",
			}
			_, err := event.Marshall(&msg)
			Expect(err).To(Succeed())
		})
		It("Publish a malformated (proto format)", func() {
			event := kafka.Event{}
			_, err := event.Marshall(make(chan int))
			Expect(err).NotTo(Succeed())
		})
	})
})
