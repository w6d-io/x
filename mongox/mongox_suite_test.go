//go:build !integration

package mongox_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMongoDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mongo Suite")
}
