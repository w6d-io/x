//go:build !integration
// +build !integration

package mongox_test

import (
	"errors"
	"time"

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
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.CreateIndexes(mongo.IndexModel{})
			Expect(err).To(Succeed())
		})
		It("index one failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.CreateIndexes(mongo.IndexModel{})
			Expect(err).NotTo(Succeed())
		})
		It("create one failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrCreateIndex: errors.New("error index"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.CreateIndexes(mongo.IndexModel{})
			Expect(err).NotTo(Succeed())
		})
		It("list index success", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ListSpecificationResult: []*mongo.IndexSpecification{
						{
							Name: "my_index",
						},
					},
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			indexes, err := client.ListIndexes()
			Expect(err).To(Succeed())
			Expect(len(indexes)).To(Equal(1))
		})
		It("list index failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			_, err := client.ListIndexes()
			Expect(err).NotTo(Succeed())
		})
		It("list index failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrListSpecifications: errors.New("error index"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			_, err := client.ListIndexes()
			Expect(err).NotTo(Succeed())
		})
		It("drop index success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.DropIndex("")
			Expect(err).To(Succeed())
		})
		It("drop index failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.DropIndex("")
			Expect(err).NotTo(Succeed())
		})
		It("drop index failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrDropOne: errors.New("error index"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.DropIndex("")
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
