package redisx_test

import (
	"context"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/redisx"
)

var _ = Describe("Queue methods", func() {
	Context("Go throw all queue method behaviour", func() {
		When("deals with RPush", func() {
			It("successfully push value into keys", func() {
				m := RedisDB{
					ClientAPI: &MockClient{},
				}
				ctx := context.Background()
				Expect(m.RPush(ctx, "test", "test")).To(Succeed())
			})
			It("should failed on connection", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				Expect(m.RPush(ctx, "test", "test")).ToNot(Succeed())

			})
			It("fails to push value into keys", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						RPushErr: errors.New("error pushing"),
					},
				}
				ctx := context.Background()
				Expect(m.RPush(ctx, "test", "test")).ToNot(Succeed())
			})
		})
		When("deals with BLPop", func() {
			It("successfully pop value into keys", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						BLPopValue: []string{"test"},
					},
				}
				ctx := context.Background()
				Expect(m.BLPop(ctx, 0, "test")).To(Equal([]string{"test"}))
			})
			It("should failed on connection", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.BLPop(ctx, 0, "test")
				Expect(err).To(HaveOccurred())

			})
			It("fails on pop value", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						BLPopErr: errors.New("error pop"),
					},
				}
				ctx := context.Background()
				_, err := m.BLPop(ctx, 0, "test")
				Expect(err).To(HaveOccurred())
			})
			It("pops an empty value", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						BLPopErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				_, err := m.BLPop(ctx, 0, "test")
				Expect(err).ToNot(HaveOccurred())
			})
		})
		When("deals with LPop", func() {
			It("successfully pop value into keys", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						LPopValue: "test",
					},
				}
				ctx := context.Background()
				Expect(m.LPop(ctx, "test")).To(Equal("test"))
			})
			It("should failed on connection", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						PingErr: errors.New("error connecting"),
					},
				}
				ctx := context.Background()
				_, err := m.LPop(ctx, "test")
				Expect(err).To(HaveOccurred())

			})
			It("fails pop value", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						LPopErr: errors.New("error pop"),
					},
				}
				ctx := context.Background()
				_, err := m.LPop(ctx, "test")
				Expect(err).To(HaveOccurred())
			})
			It("fails pop value", func() {
				m := RedisDB{
					ClientAPI: &MockClient{
						LPopErr: errors.New(NilRedis),
					},
				}
				ctx := context.Background()
				_, err := m.LPop(ctx, "test")
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
