//go:build !integration

package mongox_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/mongox"
)

var _ = Describe("Aggregate", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("aggregate success", func() {
			m := &mongox.MongoDB{
				ClientAPI:  &mongox.MockClient{},
				Collection: "collection",
			}
			client := m.SetOptions(mongox.Timeout(10 * time.Second))
			var v interface{}
			err := client.Aggregate(nil, &v)
			Expect(err).To(Succeed())
		})
		It("aggregate error", func() {
			m := &mongox.MongoDB{
				ClientAPI: &mongox.MockClient{
					ErrAggregate: errors.New("error aggregate"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(mongox.Timeout(10 * time.Second))
			var v interface{}
			err := client.Aggregate(nil, &v)
			Expect(err).NotTo(Succeed())
		})
		It("aggregate error connect", func() {
			m := &mongox.MongoDB{
				ClientAPI: &mongox.MockClient{
					ErrConnect: errors.New("error connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(mongox.Timeout(10 * time.Second))
			var v interface{}
			err := client.Aggregate(nil, &v)
			Expect(err).NotTo(Succeed())
		})
		It("aggregate error cursor", func() {
			m := &mongox.MongoDB{
				ClientAPI: &mongox.MockClient{
					ErrorCursorAll: errors.New("error cursor"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(mongox.Timeout(10 * time.Second))
			var v interface{}
			err := client.Aggregate(nil, &v)
			Expect(err).NotTo(Succeed())
		})
	})
})
