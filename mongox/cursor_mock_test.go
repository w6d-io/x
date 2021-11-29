//go:build !integration
// +build !integration

package mongox_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Cursor Mock", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			ctx := context.Background()
			cursor, err := m.GetCollection().Find(ctx, nil)
			Expect(err).To(Succeed())
			var v interface{}
			err = m.SetCursor(cursor).All(ctx, v)
			Expect(err).To(Succeed())
			_ = m.SetCursor(cursor).Next(ctx)
			err = m.SetCursor(cursor).Decode(v)
			Expect(err).To(Succeed())
		})
	})
})
