//go:build !integration

package mongox_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("CountDocuments", func() {
	Context("", func() {
		It("success to count", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					CountResult: 10,
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var count int64
			Expect(client.CountDocuments(nil, &count)).To(Succeed())
			Expect(count).To(Equal(int64(10)))
		})

		It("raises an error due to nil count", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					CountResult: 10,
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			Expect(client.CountDocuments(nil, nil)).To(HaveOccurred())
		})
		It("raises an error due to connection", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrConnect: errors.New("error connect"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var count int64
			Expect(client.CountDocuments(nil, &count)).To(HaveOccurred())
		})
		It("fails on count", func() {
			m := &MongoDB{
				ClientAPI: &MockClient{
					ErrCount: errors.New("error to count"),
				},
				Collection: "collection",
			}
			client := m.SetOptions(Timeout(10 * time.Second))
			var count int64
			Expect(client.CountDocuments(nil, &count)).To(HaveOccurred())
		})
	})
})
