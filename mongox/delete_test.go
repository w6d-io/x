//go:build !integration
// +build !integration

package mongox_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Delete", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("delete one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			err := m.Delete(nil)
			Expect(err).To(Succeed())
		})
		It("delete one failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			err := m.Delete(nil)
			Expect(err).NotTo(Succeed())
		})
		It("delete one failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrDeleteOne: errors.New("error delete one"),
				},
				Collection: "collection",
			}
			err := m.Delete(nil)
			Expect(err).NotTo(Succeed())
		})
		It("delete many success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			err := m.DeleteAll()
			Expect(err).To(Succeed())
		})
		It("delete many failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			err := m.DeleteAll()
			Expect(err).NotTo(Succeed())
		})
		It("delete many failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrDeleteMany: errors.New("error delete many"),
				},
				Collection: "collection",
			}
			err := m.DeleteAll()
			Expect(err).NotTo(Succeed())
		})
	})
})
