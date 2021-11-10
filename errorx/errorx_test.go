package errorx_test

import (
	"context"
	"net/http"
	"net/http/httptest"

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
				Cause:      errors.New("unit test"),
				StatusCode: http.StatusInternalServerError,
				Message:    "message error",
			}
			Expect(e.Error()).To(Equal(errors.Wrap(errUnitTest, "message error").Error()))
			Expect(e.GetStatusCode()).To(Equal(http.StatusInternalServerError))
			Expect(e.GetMessage()).To(Equal("message error"))
			Expect(e.GetCause().Error()).To(Equal("unit test"))
			Expect(errorx.Error{Message: "message"}.EditMessage("edit message")).To(Equal(&errorx.Error{Message: "edit message"}))
			Expect(errorx.Error{Cause: errors.New("cause")}.EditCause(errors.New("edit cause"))).To(Equal(&errorx.Error{Cause: errors.New("edit cause")}))
			Expect(errorx.Error{Code: "code"}.EditCode("edit code")).To(Equal(&errorx.Error{Code: "edit code"}))
			Expect(errorx.Error{StatusCode: 500}.EditStatusCode(500)).To(Equal(&errorx.Error{StatusCode: 500}))
			Expect(errorx.Error{}.EditStatusCode(500).EditMessage("edit message").EditCause(errors.New("edit cause")).EditCode("edit code")).To(Equal(&errorx.Error{Message: "edit message", Code: "edit code", Cause: errors.New("edit cause"), StatusCode: 500}))
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
			Expect(errorx.Error2code(&errorx.Error{StatusCode: http.StatusContinue})).To(Equal(http.StatusContinue))
			Expect(errorx.Error2code(errorx.ErrTokenNotFound)).To(Equal(http.StatusUnauthorized))
			Expect(errorx.Error2code(errorx.ErrMethod)).To(Equal(http.StatusBadRequest))
			Expect(errorx.Error2code(errors.New("internal server error test"))).To(Equal(http.StatusInternalServerError))
			w := httptest.NewRecorder()
			errorx.ErrorEncoder(nil, errorx.ErrTokenCheck, w)
			Expect(w.Code).To(Equal(http.StatusServiceUnavailable))
			w = httptest.NewRecorder()
			errorx.ErrorEncoder(nil, &errorx.Error{
				Cause:      errorx.ErrServiceUnavailable,
				StatusCode: 600,
				Code:       "unit_test_with_error",
				Message:    "this error should be raised",
			}, w)
			Expect(w.Code).To(Equal(600))
			Expect(w.Body.String()).To(Equal("{\"code\":\"unit_test_with_error\",\"message\":\"this error should be raised\"}\n"))
			errorx.NewErrorHandler().Handle(context.Background(), errors.New("unit-test"))
		})
	})
})
