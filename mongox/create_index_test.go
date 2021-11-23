// +build !integration

package mongox_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	mgoOtions "go.mongodb.org/mongo-driver/mongo/options"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("CreateIndexes", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("index one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Database:   "db",
				Collection: "collection",
			}
			err := m.CreateIndexes(mongo.IndexModel{})
			Expect(err).To(Succeed())
		})
		It("index one failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Database:   "db",
				Collection: "collection",
			}
			err := m.CreateIndexes(mongo.IndexModel{})
			Expect(err).NotTo(Succeed())
		})
		It("index one failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrIndex: errors.New("error index"),
				},
				Database:   "db",
				Collection: "collection",
			}
			err := m.CreateIndexes(mongo.IndexModel{})
			Expect(err).NotTo(Succeed())
		})
		It("get index", func() {
			clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
			clt, _ := mongo.NewClient(clientOptions)
			c := &Client{
				Client:     clt,
				Database:   "db",
				Collection: "collection",
			}
			_ = c.GetCollection().GetIndex()
		})
	})
})
