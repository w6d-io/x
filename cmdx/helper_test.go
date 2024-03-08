package cmdx_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/w6d-io/x/errorx"

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
			cmdx.Must(&errorx.Error{
				Cause:      fmt.Errorf("test raise error"),
				StatusCode: 500,
				Code:       "unit_test_code",
				Message:    "should see this error",
			}, "that's it")
			Ω(got).To(Equal(1))
		})
		It("should execute", func() {
			cmdx.Should(errors.New("unit-test"), "it isn't ok")
		})
	})
})
