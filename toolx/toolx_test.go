package toolx_test

import (
	"os"
	"sort"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/w6d-io/x/toolx"
)

var _ = Describe("Toolx functions", func() {
	Context("", func() {
		It("with InArray is true", func() {
			r := toolx.InArray("a", []string{"a", "b", "c", "d"})
			Expect(r).To(Equal(true))
		})
		It("with InArray is false", func() {
			r := toolx.InArray("z", []string{"a", "b", "c", "d"})
			Expect(r).To(Equal(false))
		})
		It("with Contains is true", func() {
			r := toolx.Contains([]string{"a", "b", "c", "d"}, "a")
			Expect(r).To(Equal(true))
		})
		It("with Contains is false", func() {
			r := toolx.Contains([]string{"a", "b", "c", "d"}, "e")
			Expect(r).To(Equal(false))
		})
		It("", func() {
			myMap := map[string]int{"a": 1, "b": 2, "c": 3}
			res := toolx.KeysMap(myMap)
			sort.Strings(res)
			Expect(res).To(Equal([]string{"a", "b", "c"}))
		})

		It("returns the default value", func() {
			v := toolx.Getenv("TEST_GO", "test")
			Expect(v).To(Equal("test"))
		})
		It("returns the default value", func() {
			v := toolx.Getenv("TEST_GO", "test")
			Expect(v).To(Equal("test"))
		})
		It("returns empty value", func() {
			v := toolx.Getenv("TEST_GO")
			Expect(v).To(Equal(""))
		})
		It("returns empty value", func() {
			os.Setenv("TEST_GO", "test_go")
			v := toolx.Getenv("TEST_GO")
			Expect(v).To(Equal("test_go"))
		})
	})
})
