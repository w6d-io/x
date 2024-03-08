//go:build !integration

package redisx_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/redisx"
)

var _ = Describe("Gets", func() {
	Context("", func() {
		When("", func() {
			It("Get fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.Get(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("Get fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						GetErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				_, err := m.Get(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("Get fails while deleting (nil error)", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						GetErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				_, err := m.Get(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("Get success while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				_, err := m.Get(ctx, "key")
				Expect(err).To(Succeed())
			})

			It("HGet fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				err := m.HGet(ctx, "key", "field", nil)
				Expect(err).NotTo(Succeed())
			})
			It("HGet fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				err := m.HGet(ctx, "key", "field", nil)
				Expect(err).NotTo(Succeed())
			})
			It("HGet fails while getting (nil error)", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				err := m.HGet(ctx, "key", "field", nil)
				Expect(err).NotTo(Succeed())
			})
			It("HGet success while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetValue: `{"payload": "hello world"}`,
					},
				}
				ctx := context.Background()
				var ret interface{}
				err := m.HGet(ctx, "key", "field", &ret)
				Expect(err).To(Succeed())
			})
			It("HGet failure unmarshal while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetValue: "hello world",
					},
				}
				ctx := context.Background()
				var ret interface{}
				err := m.HGet(ctx, "key", "field", &ret)
				Expect(err).NotTo(Succeed())
			})

			It("HGetAll fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.HGetAll(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("HGetAll fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetAllErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				_, err := m.HGetAll(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("HGetAll fails while deleting (nil error)", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetAllErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				_, err := m.HGetAll(ctx, "key")
				Expect(err).NotTo(Succeed())
			})
			It("HGetAll success while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						HGetAllValue: map[string]string{},
					},
				}
				ctx := context.Background()
				_, err := m.HGetAll(ctx, "key")
				Expect(err).To(Succeed())
			})

			It("Keys fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.Keys(ctx, "pattern")
				Expect(err).NotTo(Succeed())
			})
			It("Keys fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						KeysErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				_, err := m.Keys(ctx, "pattern")
				Expect(err).NotTo(Succeed())
			})
			It("Keys fails while getting (nil error)", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						KeysErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				_, err := m.Keys(ctx, "pattern")
				Expect(err).NotTo(Succeed())
			})
			It("Keys success while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						KeysValue: []string{},
					},
				}
				ctx := context.Background()
				_, err := m.Keys(ctx, "pattern")
				Expect(err).To(Succeed())
			})

			It("Scan fails while connecting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.Scan(ctx, "match", 0)
				Expect(err).NotTo(Succeed())
			})
			It("Scan fails while getting", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						ScanErr: errors.New("error getting"),
					},
				}
				ctx := context.Background()
				_, err := m.Scan(ctx, "match", 0)
				Expect(err).NotTo(Succeed())
			})
			It("Scan empty keys", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						ScanValue: []string{"key0"},
					},
				}
				ctx := context.Background()
				_, err := m.Scan(ctx, "match", 0)
				Expect(err).To(Succeed())
			})
		})
	})
})
