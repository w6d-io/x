package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/pkg/errors"

	"github.com/w6d-io/x/errorx"

	. "github.com/onsi/ginkgo"
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
			ctx := httpx.BeforeHttpFunc(context.Background(), req)
			Expect(ctx.Value("correlation_id")).ShouldNot(BeNil())
			Expect(ctx.Value("kind")).ShouldNot(BeNil())
			Expect(ctx.Value("ipaddress")).ShouldNot(BeNil())
		})
	})
	Context("On bad request", func() {
		It("does not fill ipaddress because of bad request ip address", func() {
			req := &http.Request{
				RemoteAddr: "[10.0.1.1:4242",
			}
			ctx := httpx.BeforeHttpFunc(context.Background(), req)
			Expect(ctx.Value("ipaddress")).Should(BeEmpty())
		})
		It("fills ip address with an hyphen", func() {
			req := &http.Request{
				RemoteAddr: "0.1.1:4242",
			}
			ctx := httpx.BeforeHttpFunc(context.Background(), req)
			Expect(ctx.Value("ipaddress")).Should(Equal("-"))
		})
	})
	Context("Encode http response", func() {
		It("return nil and write endpoint.Failer into http.ResponseWriter", func() {
			w := httptest.NewRecorder()
			rsp := &failedResponse{
				err: errors.New("test"),
			}
			err := httpx.EncodeHTTPResponse(w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: &errorx.Error{Code: 800},
			}
			err = httpx.EncodeHTTPResponse(w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: errorx.ErrTokenNotFound,
			}
			err = httpx.EncodeHTTPResponse(w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: errorx.ErrMethod,
			}
			err = httpx.EncodeHTTPResponse(w, rsp)
			Expect(err).ToNot(HaveOccurred())

			rsp = &failedResponse{
				err: errorx.ErrTokenCheck,
			}
			err = httpx.EncodeHTTPResponse(w, rsp)
			Expect(err).ToNot(HaveOccurred())
		})
		It("json encodes the error into http.ResponseWriter", func() {
			w := httptest.NewRecorder()
			err := httpx.EncodeHTTPResponse(w, struct {
				Error string
			}{
				Error: "failed",
			})
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
