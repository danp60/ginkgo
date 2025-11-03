package C_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
