package logx_test

import (
	"context"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/logx"
)

var _ = Describe("in logx package", func() {
	Context("get values from context", func() {
		It("get all elements", func() {
			By("create variable")
			correlationID := uuid.New().String()

			By("creating the context")
			ctx := context.WithValue(context.Background(), "correlation_id", correlationID)
			ctx = context.WithValue(ctx, "kind", "test")
			ctx = context.WithValue(ctx, "ipaddress", "127.0.0.42")

			values := logx.GetLogValues(ctx)
			Expect(values).To(Not(BeNil()))
			Expect(values[1].(string)).To(Equal(correlationID))
			Expect(values[3].(string)).To(Equal("test"))
			Expect(values[5].(string)).To(Equal("127.0.0.42"))
			Expect(logx.GetLogValues(nil)).To(BeEmpty())

			Expect(logx.WithName(ctx, "UnitTest")).ToNot(BeNil())
			Expect(logx.GetCorrelationID(ctx)).To(Equal(correlationID))
			Expect(logx.GetCorrelationID(nil)).To(BeEmpty())
			Expect(logx.GetCorrelationID(context.Background())).To(BeEmpty())
		})
	})
})
