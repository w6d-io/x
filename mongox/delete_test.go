package mongox_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Delete", func() {
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
			filter := bson.D{primitive.E{Key: "test", Value: "test"}}
			ctx := context.Background()
			_, err := c.DeleteOne(ctx, "", "", filter)
			Expect(err).ToNot(Succeed())
		})
	})
})
