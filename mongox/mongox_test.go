package mongox_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Mongo", func() {
	Context("", func() {
		When("client is set", func() {
			It("get client created before", func() {
				_, err := GetClient()
				Expect(err).To(Succeed())
				Expect(GetClient()).ToNot(BeNil())
			})
		})
	})
})
