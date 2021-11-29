package cmdx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/w6d-io/x/cmdx"
)

var _ = Describe("helper functions testing", func() {
	Context("checking Must behaviour", func() {
		It("Must works without printing", func() {
			cmdx.Must(nil, "never write")
		})
		It("Must works without printing", func() {
			var got int
			myExit := func(code int) {
				got = code
			}
			cmdx.OsExit = myExit
			cmdx.Must(errors.New("test exits"), "all is good")
			Ω(got).To(Equal(1))
		})
		It("should execute", func() {
			cmdx.Should("it isn't ok", errors.New("unit-test"))
		})
	})
})
