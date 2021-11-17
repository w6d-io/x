package mongox_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("CreateIndexes", func() {
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
			var expire int32 = 30
			keys := bsonx.Doc{{Key: "expirationTime", Value: bsonx.Int32(int32(1))}}
			data := mongo.IndexModel{
				Keys: keys,
				Options: &options.IndexOptions{
					ExpireAfterSeconds: &expire,
				},
			}
			ctx := context.Background()
			err := c.CreateOne(ctx, "", "", data)
			Expect(err).ToNot(Succeed())
		})
	})
})
