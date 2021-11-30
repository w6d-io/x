package grpcx_test

import (
	"context"
	"github.com/w6d-io/x/logx"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"google.golang.org/grpc/peer"

	"github.com/w6d-io/x/grpcx"
)

var _ = Describe("run all grpc package functions", func() {
	Context("in the main file", func() {
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("gets options", func() {
			o := grpcx.GrpcOptions()
			Expect(len(o)).To(Equal(2))
		})
		It("gets context field with all metadata", func() {
			var ctx context.Context
			By("Set peer into context", func() {
				addr := &mockAddr{Address: "127.0.0.1:8080"}
				ctx = context.Background()
				p := &peer.Peer{
					Addr: addr,
				}
				ctx = peer.NewContext(ctx, p)
			})
			nCtx := grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value(logx.CorrelationId)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind).(string)).To(Equal("grpc"))
			Expect(nCtx.Value(logx.IpAddress)).ToNot(BeNil())
			Expect(nCtx.Value(logx.IpAddress).(string)).To(Equal("127.0.0.1"))
		})
		It("gets context field without ipaddress into metadata", func() {
			var ctx context.Context
			By("Set a bad host port into peer context", func() {
				addr := &mockAddr{Address: "127.0.0.1#8080"}
				ctx = context.Background()
				p := &peer.Peer{
					Addr: addr,
				}
				ctx = peer.NewContext(ctx, p)
			})
			nCtx := grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value(logx.CorrelationId)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind).(string)).To(Equal("grpc"))
			Expect(nCtx.Value(logx.IpAddress)).To(BeNil())

			By("set a bad ip address", func() {
				addr := &mockAddr{Address: "127.0.0..1:8080"}
				ctx = context.Background()
				p := &peer.Peer{
					Addr: addr,
				}
				ctx = peer.NewContext(ctx, p)

			})
			nCtx = grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value(logx.CorrelationId)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind).(string)).To(Equal("grpc"))
			Expect(nCtx.Value(logx.IpAddress)).ToNot(BeNil())
			Expect(nCtx.Value(logx.IpAddress).(string)).To(Equal("-"))
			By("an empty context")
			ctx = context.Background()
			nCtx = grpcx.BeforeGrpcFunc(ctx, nil)
			Expect(nCtx.Value(logx.CorrelationId)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind)).ToNot(BeNil())
			Expect(nCtx.Value(logx.Kind).(string)).To(Equal("grpc"))
			Expect(nCtx.Value(logx.IpAddress)).To(BeNil())
		})
	})
})
