package k8x_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/k8x"
)

var _ = Describe("", func() {
	Context("", func() {
		It("get as new k8s instance", func() {
			Expect(k8x.New()).ToNot(BeNil())
		})
		It("get a kubernetes client", func() {
			k := k8x.New()
			k.SetConfig(cfg)
			c, err := k.GetClient(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).ToNot(BeNil())
		})
		It("gets the client already recorded", func() {
			k := k8x.New()
			k.SetClient(k8sClient)
			c, err := k.GetClient(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(fmt.Sprintf("%p", c)).To(Equal(fmt.Sprintf("%p", k8sClient)))
		})
		It("gets a new config", func() {
			k := k8x.New()
			c, err := k.GetConfig(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).ToNot(BeNil())
		})
		It("fails to get client due to config error on config path", func() {
			k := k8x.New(k8x.Context("/nopath"))
			_, err := k.GetClient(ctx)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("get namespace failed:"))
		})
		//It("fails to get client due to config error on ", func() {
		//    k := k8x.NewWithConfig("/nopath", "", "")
		//    _, err := k.GetClient(ctx)
		//    Expect(err).To(HaveOccurred())
		//    Expect(err.Error()).To(Equal("get namespace failed: stat /nopath: no such file or directory"))
		//})
		It("fails to get client due to config error on ", func() {
			k := k8x.New(k8x.ConfigPath("testdata/test1.yaml"), k8x.Context("test-1"), k8x.Namespace("default"))
			_, err := k.GetClient(ctx)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("get namespace failed: Get \"https://localhost:6444/api?timeout=32s\": dial tcp [::1]:6444: connect: connection refused"))
		})
	})
})
