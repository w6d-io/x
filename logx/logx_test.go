package logx_test

import (
	"context"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/logx"
)

var _ = Describe("in logx package", func() {
	Context("get values from context", func() {
		It("get all elements", func() {
			By("create variable")
			correlationID := uuid.New().String()

			By("creating the context")
			ctx := context.WithValue(context.Background(), logx.CorrelationID, correlationID)
			ctx = context.WithValue(ctx, logx.Kind, "test")
			ctx = context.WithValue(ctx, logx.IPAddress, "127.0.0.42")
			m := map[string]string{
				"kind":           "test",
				"ipaddress":      "127.0.0.42",
				"correlation_id": correlationID,
			}
			values := logx.GetLogValues(ctx)
			Expect(values).To(Not(BeNil()))
			for n := 0; n < len(values); n = n + 2 {
				Expect(values[n+1].(string)).To(Equal(m[values[n].(string)]))
			}
			//Expect(values[1].(string)).To(Equal(correlationID))
			//Expect(values[3].(string)).To(Equal("test"))
			//Expect(values[5].(string)).To(Equal("127.0.0.42"))
			Expect(logx.GetLogValues(context.TODO())).To(BeEmpty())

			Expect(logx.WithName(ctx, "UnitTest")).ToNot(BeNil())
			Expect(logx.GetCorrelationID(ctx)).To(Equal(correlationID))
			Expect(logx.GetCorrelationID(context.TODO())).To(BeEmpty())
			Expect(logx.GetCorrelationID(context.Background())).To(BeEmpty())
		})
	})
})
