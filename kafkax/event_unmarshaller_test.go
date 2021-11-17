package kafkax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("Event Unmarshaller", func() {
	Context("", func() {
		var (
			schema kafka.Schema
		)
		BeforeEach(func() {
			schema = kafka.Schema{
				Id:     1,
				Schema: "syntax = \"proto3\";\npackage kafka;\n\nmessage Test {\n string Message= 1;\n}\n",
			}
		})
		AfterEach(func() {
		})
		It("Marshall then Unmarshall (json format)", func() {
			event := kafka.Event{}
			msg := TTest{
				Message: "Hello World",
			}
			value, err := event.Marshall(msg)
			Expect(err).To(Succeed())
			event.Value = value
			d := make(map[string]interface{})
			err = event.Unmarshal(d)
			Expect(err).To(Succeed())
			Expect(d["Message"]).To(Equal("Hello World"))
		})
		It("Marshall then Unmarshall (proto format)", func() {
			msg := Test{
				Message: "Hello World",
			}
			event := kafka.Event{
				Topic: "kafka",
				SchemaRegistry: &kafka.MockSchemaRegistry{
					Schema: schema,
				},
			}
			value, err := event.Marshall(&msg)
			event.Value = value
			Expect(err).To(Succeed())
			d := make(map[string]interface{})
			err = event.Unmarshal(d)
			Expect(err).To(Succeed())
			Expect(d["Message"]).To(Equal(msg.Message))
		})
	})
})
