package toolx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/toolx"
)

var _ = Describe("InArray", func() {
	Context("", func() {
		It("is true", func() {
			r := toolx.InArray("a", []string{"a", "b", "c", "d"})
			Expect(r).To(Equal(true))
		})
		It("is true", func() {
			r := toolx.InArray("z", []string{"a", "b", "c", "d"})
			Expect(r).To(Equal(false))
		})
	})
})
