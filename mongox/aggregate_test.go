//go:build !integration
// +build !integration

package mongox_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Aggregate", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("aggregate success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			var v interface{}
			err := m.Aggregate(nil, &v)
			Expect(err).To(Succeed())
		})
		It("aggregate error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrAggregate: errors.New("error aggregate"),
				},
				Collection: "collection",
			}
			var v interface{}
			err := m.Aggregate(nil, &v)
			Expect(err).NotTo(Succeed())
		})
		It("aggregate error connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error connect"),
				},
				Collection: "collection",
			}
			var v interface{}
			err := m.Aggregate(nil, &v)
			Expect(err).NotTo(Succeed())
		})
		It("aggregate error cursor", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrorCursorAll: errors.New("error cursor"),
				},
				Collection: "collection",
			}
			var v interface{}
			err := m.Aggregate(nil, &v)
			Expect(err).NotTo(Succeed())
		})
	})
})
