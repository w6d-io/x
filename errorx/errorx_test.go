package errorx_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/pkg/errors"
    "github.com/w6d-io/x/errorx"
    "net/http"
)

var _ = Describe("Error x", func() {
    Context("Method", func() {
        It("check", func() {
            By("build errorx")
            e := errorx.Error{
                Cause: errors.New("unit test"),
                Code: http.StatusInternalServerError,
                Message: "message error",
            }
            Expect(e.Error()).To(Equal(errors.Wrap(errors.New("unit test"), "message error").Error()))
            Expect(e.GetStatusCode()).To(Equal(http.StatusInternalServerError))
            Expect(e.GetMessage()).To(Equal("message error"))
            By("remove cause")
            e.Cause = nil
            Expect(e.Error()).To(Equal("message error"))
            //Expect(e.GetMessage()).To(Equal("message error"))

        })
    })
})
