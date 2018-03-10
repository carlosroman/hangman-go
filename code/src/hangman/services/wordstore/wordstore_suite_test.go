package wordstore_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWordstore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wordstore Suite")
}
