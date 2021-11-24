// +build !integration

package mongox_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("update", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {
		})
		AfterEach(func() {
		})
		It("update one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			err := m.Update(nil, nil)
			Expect(err).To(Succeed())
		})
		It("update one connect error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("connect error"),
				},
				Collection: "collection",
			}
			err := m.Update(nil, nil)
			Expect(err).NotTo(Succeed())
		})
		It("update one error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrUpdateOne: errors.New("update one error"),
				},
				Collection: "collection",
			}
			err := m.Update(nil, nil)
			Expect(err).NotTo(Succeed())
		})
		It("upsert one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			err := m.Upsert(nil, nil)
			Expect(err).To(Succeed())
		})
		It("upsert one connect error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("connect error"),
				},
				Collection: "collection",
			}
			err := m.Upsert(nil, nil)
			Expect(err).NotTo(Succeed())
		})
		It("upsert one error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrReplaceOne: errors.New("update one error"),
				},
				Collection: "collection",
			}
			err := m.Upsert(nil, nil)
			Expect(err).NotTo(Succeed())
		})
		It("findoneandupdate one success", func() {
			m := &MongoDB{
				ClientAPI:  &MockClient{},
				Collection: "collection",
			}
			err := m.FindAndUpdate(nil, nil, nil)
			Expect(err).To(Succeed())
		})
		It("findoneandupdate one connect error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("connect error"),
				},
				Collection: "collection",
			}
			err := m.FindAndUpdate(nil, nil, nil)
			Expect(err).NotTo(Succeed())
		})
		It("findoneandupdate one error", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrorSingleResultDecode: errors.New("single result error"),
				},
				Collection: "collection",
			}
			err := m.FindAndUpdate(nil, nil, nil)
			Expect(err).NotTo(Succeed())
		})
	})
})
