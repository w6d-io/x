//go:build !integration

package redisx_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/redisx"
)

var _ = Describe("Dels", func() {
	Context("", func() {
		When("", func() {
			It("fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				err := m.HDel(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("fails while deleting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HDelErr: errors.New("error deleting"),
					},
				}
				ctx := context.Background()
				err := m.HDel(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("fails while deleting (nil error)", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HDelErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				err := m.HDel(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("success while deleting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				err := m.HDel(ctx, "key")
				Expect(err).To(Succeed())
			})
		})
	})
})
