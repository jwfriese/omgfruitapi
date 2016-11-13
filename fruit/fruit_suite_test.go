package fruit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFruit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fruit Suite")
}
