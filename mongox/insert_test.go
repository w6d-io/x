package mongox_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Insert", func() {
	Context("", func() {
		var (
			c   *Client
			err error
		)
		BeforeEach(func() {
			c, err = GetClient()
			Expect(err).To(Succeed())
		})
		AfterEach(func() {
			c = nil
			err = nil
		})
		It("", func() {
			type user struct {
				Name    string
				Project string
			}
			data := []user{
				{
					Name:    "test",
					Project: "test",
				},
			}
			ctx := context.Background()
			_, err := c.InsertOne(ctx, "", "", data)
			Expect(err).ToNot(Succeed())
		})
	})
})
