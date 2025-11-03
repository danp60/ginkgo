package hanging_suite_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHangingSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HangingSuite Suite")
}
