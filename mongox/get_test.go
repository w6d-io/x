//go:build !integration

package mongox_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Get", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("get one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var v interface{}
			err := client.Get(nil, &v)
			Expect(err).To(Succeed())
		})
		It("get one error find", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrFind: errors.New("error find"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var v interface{}
			err := client.Get(nil, &v)
			Expect(err).NotTo(Succeed())
		})
		It("get one error connect", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var v interface{}
			err := client.Get(nil, &v)
			Expect(err).NotTo(Succeed())
		})
		It("get one error cursor", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrorCursorAll: errors.New("error cursor"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var v interface{}
			err := client.Get(nil, &v)
			Expect(err).NotTo(Succeed())
		})
	})
})
