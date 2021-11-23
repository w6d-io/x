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

var _ = Describe("Client", func() {
	Context("", func() {
		When("Create client", func() {
			It("fails option", func() {
				cfg := &Mongo{
					URL:        "",
					AuthSource: "db",
				}
				_, err := cfg.New("", WithProtoCodec())
				Expect(err).NotTo(Succeed())
			})
			It("success option", func() {
				cfg := &Mongo{
					URL:        "mongodb://127.0.0.1",
					AuthSource: "db",
				}
				_, err := cfg.New("", WithProtoCodec())
				Expect(err).To(Succeed())
			})
			It("success option, get collection", func() {
				cfg := &Mongo{
					URL:        "mongodb://127.0.0.1",
					AuthSource: "db",
				}
				_, err := cfg.New("", WithProtoCodec())
				Expect(err).To(Succeed())
			})
			It("failure connect", func() {
				m := MongoDB{
					ClientAPI: &MockClient{
						ErrConnect: errors.New("error while connecting"),
					},
				}
				err := m.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("success connect", func() {
				m := MongoDB{
					ClientAPI: &MockClient{},
				}
				err := m.Connect()
				Expect(err).To(Succeed())
				m.GetCollection()
			})
			It("failure ping", func() {
				m := MongoDB{
					ClientAPI: &MockClient{
						ErrPing: errors.New("error while ping"),
					},
				}
				err := m.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("get collection", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				_ = c.GetCollection()
			})
			It("set cursor", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				_ = c.SetCursor(nil)
			})
			It("set single result", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				_ = c.SetSingleResult(nil)
			})
		})
	})
})
