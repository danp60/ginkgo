package example_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Suite")
}
