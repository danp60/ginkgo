package D_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}
