//go:build !integration

package mongox_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
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
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.Delete(nil)
			Expect(err).To(Succeed())
		})
		It("delete one failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.Delete(nil)
			Expect(err).NotTo(Succeed())
		})
		It("delete one failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrDeleteOne: errors.New("error delete one"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.Delete(nil)
			Expect(err).NotTo(Succeed())
		})
		It("delete many success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.DeleteAll()
			Expect(err).To(Succeed())
		})
		It("delete many failure on connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error on connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.DeleteAll()
			Expect(err).NotTo(Succeed())
		})
		It("delete many failure", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrDeleteMany: errors.New("error delete many"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			err := client.DeleteAll()
			Expect(err).NotTo(Succeed())
		})
	})
})
