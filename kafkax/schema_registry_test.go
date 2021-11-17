package kafkax_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kafka "github.com/w6d-io/x/kafkax"
)

var _ = Describe("SchemaRegistry", func() {
	Context("", func() {
		var (
			schema kafka.Schema
		)
		BeforeEach(func() {
			schema = kafka.Schema{
				Id:     1,
				Schema: "",
			}
		})
		AfterEach(func() {
		})
		It("GetLatestSchema", func() {
			client := kafka.MockSchemaRegistry{
				Schema: schema,
			}
			schema, err := client.GetLatestSchema("test")
			Expect(err).To(Succeed())
			Expect(schema.Id).To(Equal(1))
			Expect(schema.Schema).To(Equal(""))
		})
	})
})
