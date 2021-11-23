// +build !integration

package mongox_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Counter", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("incr success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Database:   "db",
				Collection: "collection",
			}
			_, err := m.Incr("key")
			Expect(err).To(Succeed())
		})
		It("incr error connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error connect"),
				},
				Database:   "db",
				Collection: "collection",
			}
			_, err := m.Incr("key")
			Expect(err).NotTo(Succeed())
		})
		It("incr error find", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrFind: errors.New("error find"),
				},
				Database:   "db",
				Collection: "collection",
			}
			_, err := m.Incr("key")
			Expect(err).NotTo(Succeed())
		})
		It("incr error insert", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrInsertOne: errors.New("error insert"),
				},
				Database:   "db",
				Collection: "collection",
			}
			_, err := m.Incr("key")
			Expect(err).NotTo(Succeed())
		})
		It("incr error update", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrUpdateOne: errors.New("error update"),
				},
				Database:   "db",
				Collection: "collection",
			}
			_, err := m.Incr("key")
			Expect(err).NotTo(Succeed())
		})
	})
})