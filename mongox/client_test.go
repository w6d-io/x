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

	"github.com/w6d-io/x/mongox"
)

var _ = Describe("Client", func() {
	Context("", func() {
		When("Create client", func() {
			It("fails option", func() {
				cfg := &mongox.Mongo{
					URL:        "",
					AuthSource: "db",
				}
				m := cfg.New()
				err := m.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("success option", func() {
				cfg := &mongox.Mongo{
					URL:        "mongodb://127.0.0.1",
					AuthSource: "db",
				}
				m := cfg.New()
				m = m.SetCollection("test")
				m = m.SetOptions(mongox.WithProtoCodec(), mongox.WithStrCodec(), mongox.Timeout(10*time.Second))
				err := m.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("failure connect", func() {
				m := &mongox.MongoDB{
					ClientAPI: &mongox.MockClient{
						ErrConnect: errors.New("error while connecting"),
					},
				}
				client := m.SetOptions(mongox.Timeout(10 * time.Second))
				err := client.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("success connect", func() {
				m := &mongox.MongoDB{
					ClientAPI: &mongox.MockClient{},
				}
				client := m.SetOptions(mongox.Timeout(10 * time.Second))
				err := client.Connect()
				Expect(err).To(Succeed())
				m.GetCollection()
			})
			It("failure ping", func() {
				m := &mongox.MongoDB{
					ClientAPI: &mongox.MockClient{
						ErrPing: errors.New("error while ping"),
					},
				}
				client := m.SetOptions(mongox.Timeout(10 * time.Second))
				err := client.Connect()
				Expect(err).NotTo(Succeed())
			})
			It("set collection", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &mongox.Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				c.SetCollection("new_collection")
			})
			It("re-set collection", func() {
				m := &mongox.MongoDB{
					ClientAPI: &mongox.MockClient{},
				}
				m.SetCollection("collection")
			})
			It("get collection", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &mongox.Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				_ = c.GetCollection()
			})
			It("set cursor", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &mongox.Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				_ = c.SetCursor(nil)
			})
			It("set single result", func() {
				clientOptions := mgoOtions.Client().ApplyURI("mongodb://127.0.0.1")
				clt, _ := mongo.NewClient(clientOptions)
				c := &mongox.Client{
					Client:     clt,
					Database:   "db",
					Collection: "collection",
				}
				_ = c.SetSingleResult(nil)
			})
		})
	})
})
