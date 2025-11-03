package A_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}
