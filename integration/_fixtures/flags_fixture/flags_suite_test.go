package flags_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFlags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flags Suite")
}
