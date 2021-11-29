//go:build !integration
// +build !integration

package mongox_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Insert", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("insert one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
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
			err := m.Insert(data)
			Expect(err).To(Succeed())
		})
		It("insert one connect failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
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
			err := m.Insert(data)
			Expect(err).NotTo(Succeed())
		})
		It("insert one failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrInsertOne: errors.New("insert one failure"),
				},
				Collection: "collection",
			}
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
			err := m.Insert(data)
			Expect(err).NotTo(Succeed())
		})
		It("insert bulk success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			var operations []*mongo.UpdateOneModel
			operations = append(operations, BuildBulkOperation(nil, nil, nil))
			err := m.InsertBulk(operations)
			Expect(err).To(Succeed())
		})
		It("insert bulk connect error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			var operations []*mongo.UpdateOneModel
			operations = append(operations, BuildBulkOperation(nil, nil, nil))
			err := m.InsertBulk(operations)
			Expect(err).NotTo(Succeed())
		})
		It("insert bulk write error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrBulkWrite: errors.New("error on bulk write"),
				},
				Collection: "collection",
			}
			var operations []*mongo.UpdateOneModel
			operations = append(operations, BuildBulkOperation(nil, nil, nil))
			err := m.InsertBulk(operations)
			Expect(err).NotTo(Succeed())
		})
	})
})
