//go:build !integration

package redisx_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/redisx"
)

var _ = Describe("Incr", func() {
	Context("", func() {
		When("", func() {
			It("Incr fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.Incr(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("Incr fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						IncrErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				_, err := m.Incr(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("Incr fails while getting (nil error)", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						IncrErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				_, err := m.Incr(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("Incr success while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				_, err := m.Incr(ctx, "key")
				Expect(err).To(Succeed())
			})
		})
	})
})
