package suite1_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAfterRunHook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Build reporting suite 1")
}
