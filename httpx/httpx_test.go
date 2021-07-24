package httpx_test

import (
	"net/http"

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
})
