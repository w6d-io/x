//go:build !integration

package redisx_test

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/redisx"
)

var _ = Describe("Sets", func() {
	Context("", func() {
		When("", func() {
			It("Set fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				err := m.Set(ctx, "key", 1*time.Second, nil)
				Expect(err).NotTo(Succeed())
			})
			It("Set fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						SetErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				err := m.Set(ctx, "key", 1*time.Second, nil)
				Expect(err).NotTo(Succeed())
			})
			It("Set success while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				err := m.Set(ctx, "key", 1*time.Second, nil)
				Expect(err).To(Succeed())
			})

			It("HSet fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				err := m.HSet(ctx, "key", "field", nil)
				Expect(err).NotTo(Succeed())
			})
			It("HSet fails while marshalling", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				err := m.HSet(ctx, "key", "field", make(chan int))
				Expect(err).NotTo(Succeed())
			})
			It("HSet fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HSetErr: errors.New("hset error"),
					},
				}
				ctx := context.Background()
				data := `{"payload": "hello world"}`
				err := m.HSet(ctx, "key", "field", data)
				Expect(err).NotTo(Succeed())
			})
			It("HSet success", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				data := `{"payload": "hello world"}`
				err := m.HSet(ctx, "key", "field", data)
				Expect(err).To(Succeed())
			})
		})
	})
})
