package errorx_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/w6d-io/x/errorx"
)

var _ = Describe("Error x", func() {
	Context("Method", func() {
		It("check", func() {
			var errUnitTest = errors.New("unit test")
			By("build errorx")
			e := errorx.Error{
				Cause:   errors.New("unit test"),
				Code:    http.StatusInternalServerError,
				Message: "message error",
			}
			Expect(e.Error()).To(Equal(errors.Wrap(errUnitTest, "message error").Error()))
			Expect(e.GetStatusCode()).To(Equal(http.StatusInternalServerError))
			Expect(e.GetMessage()).To(Equal("message error"))
			Expect(e.GetCause().Error()).To(Equal("unit test"))
			e.ShowStack()
			By("remove cause")
			e.Cause = nil
			Expect(e.Error()).To(Equal("message error"))
			Expect(errorx.New(nil, "")).To(HaveOccurred())
			Expect(errorx.NewHTTP(nil, 200, "")).To(HaveOccurred())
			err := errorx.Wrap(errUnitTest, "error wrapped")
			Expect(err).To(HaveOccurred())

			Expect(errorx.GetError(e.Cause)).ToNot(HaveOccurred())
			Expect(errorx.GetError(&e)).To(HaveOccurred())
			err = errorx.GetError(errors.New("not implement Error"))
			_, ok := err.(*errorx.Error)
			Expect(ok).To(Equal(true))
			Expect(err.(*errorx.Error).Error()).To(Equal("not implement Error"))
		})
	})
})