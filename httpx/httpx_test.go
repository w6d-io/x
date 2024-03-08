package httpx_test

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/w6d-io/x/logx"

	"github.com/pkg/errors"

	"github.com/w6d-io/x/errorx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/httpx"
)

var _ = Describe("", func() {
	Context("read user ip", func() {
		It("", func() {
			By("build http request")
			r := &http.Request{
				Header:     http.Header{},
				RemoteAddr: "127.0.0.1",
			}

			By("get user ip")
			Expect(httpx.ReadRemoteIP(r)).To(Equal("127.0.0.1"))
		})
	})
	Context("Request have an ip address", func() {
		It("add metadata in the context and return the new context", func() {
			req := &http.Request{
				RemoteAddr: "10.0.1.1:4242",
			}
			ctx := httpx.BeforeHTTPFunc(context.Background(), req)
			Expect(ctx.Value(logx.CorrelationID)).ShouldNot(BeNil())
			Expect(ctx.Value(logx.Kind)).ShouldNot(BeNil())
			Expect(ctx.Value(logx.IPAddress)).ShouldNot(BeNil())
		})
	})
	Context("On bad request", func() {
		It("does not fill ipaddress because of bad request ip address", func() {
			req := &http.Request{
				RemoteAddr: "[10.0.1.1:4242",
			}
			ctx := httpx.BeforeHTTPFunc(context.Background(), req)
			Expect(ctx.Value(logx.IPAddress)).Should(BeNil())
		})
		It("fills ip address with an hyphen", func() {
			req := &http.Request{
				RemoteAddr: "0.1.1:4242",
				URL: &url.URL{
					Path: "/test",
				},
			}
			ctx := httpx.BeforeHTTPFunc(context.Background(), req)
			Expect(ctx.Value(logx.IPAddress).(string)).Should(Equal("-"))
		})
	})
	Context("Encode http response", func() {
		It("return nil and write endpoint.Failer into http.ResponseWriter", func() {
			w := httptest.NewRecorder()
			rsp := &failedResponse{
				err: errors.New("test"),
			}
			err := httpx.EncodeHTTPResponse(ctx, w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: &errorx.Error{StatusCode: 800},
			}
			err = httpx.EncodeHTTPResponse(ctx, w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: errorx.ErrTokenNotFound,
			}
			err = httpx.EncodeHTTPResponse(ctx, w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: errorx.ErrMethod,
			}
			err = httpx.EncodeHTTPResponse(ctx, w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: errorx.ErrTokenCheck,
			}
			err = httpx.EncodeHTTPResponse(ctx, w, rsp)
			Expect(err).ToNot(HaveOccurred())

		})
		It("deals with proto response", func() {
			rsp := &Test1{
				Message: "rocks",
			}
			w := httptest.NewRecorder()
			Expect(httpx.EncodeHTTPResponse(ctx, w, rsp)).To(Succeed())
		})
		It("json encodes the error into http.ResponseWriter", func() {
			w := httptest.NewRecorder()
			err := httpx.EncodeHTTPResponse(ctx, w, struct {
				Error string
			}{
				Error: "failed",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("{\"Error\":\"failed\"}\n"))
		})
		It("json encodes the errorx.Error into http.ResponseWriter", func() {
			w := httptest.NewRecorder()
			e := errorx.Error{
				Cause:      errorx.ErrServiceUnavailable,
				StatusCode: 601,
				Code:       "httpx_write_errorx_error",
				Message:    "should write this error",
			}
			Expect(httpx.EncodeHTTPResponse(ctx, w, e)).To(Succeed())
			Expect(w.Code).To(Equal(601))
			Expect(w.Body.String()).To(Equal("{\"code\":\"httpx_write_errorx_error\",\"message\":\"should write this error\"}\n"))
		})
	})
	Context("Miscellaneous", func() {
		It("deals on not found handler", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "http://localhost", nil)
			httpx.NotFoundHandler(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
		It("tests close listener", func() {
			listener, err := net.Listen("tcp", ":0")
			Expect(err).ShouldNot(HaveOccurred())
			c := httpx.CloseListener(listener)
			c(err)
		})
	})
})
