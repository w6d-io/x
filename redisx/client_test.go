//go:build !integration

package redisx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/redisx"
)

var _ = Describe("Client", func() {
	Context("", func() {
		When("Create client", func() {
			It("fails option", func() {
				cfg := &Redis{
					AddressSvc: "",
					DB:         1,
				}
				m := cfg.New()
				err := m.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("Success", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				err := m.Connect()
				Expect(err).To(Succeed())
			})
		})
	})
})
