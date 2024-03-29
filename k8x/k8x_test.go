package k8x_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
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
			Expect(err.Error()).To(Equal("get namespace failed: exec plugin: invalid apiVersion \"client.authentication.k8s.io/v1beta\""))
		})
		Context("GetObjectContain function", func() {
			It("converts the kubernetes manifests", func() {
				obj := &testObject{
					Test: "should ok",
				}
				Expect(k8x.GetObjectContain(obj)).To(Equal("Test: should ok\n"))
			})
			It("returns a string error", func() {
				Expect(k8x.GetObjectContain(nil)).To(Equal("<ERROR>\n"))
			})
		})
	})
})
