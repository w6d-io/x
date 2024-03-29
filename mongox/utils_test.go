//go:build !integration

package mongox_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("utils", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("get level with data", func() {
			level := GetLogLevel("hello world")
			Expect(level).To(Equal(2))
		})
		It("get level without data", func() {
			level := GetLogLevel(nil)
			Expect(level).To(Equal(3))
		})
	})
})
