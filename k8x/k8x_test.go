package k8x_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/k8x"
)

var _ = Describe("Kubernetes client testing", func() {
	Context("Go throw all case", func() {
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
			k := k8x.New(k8x.ConfigPath("testdata/test1.yaml"), k8x.Context("test-1"), k8x.Namespace("default"), k8x.Timeout("5s"))
			_, err := k.GetClient(ctx)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid configuration: [command must be specified for test-1 to use exec authentication plugin, apiVersion must be specified for test-1 to use exec authentication plugin, interactiveMode must be specified for test-1 to use exec authentication plugin]"))
			//Expect(err.Error()).To(Equal("get namespace failed: Get \"https://localhost:6444/api?timeout=5s\": dial tcp [::1]:6444: connect: connection refused"))
		})
	})
})
