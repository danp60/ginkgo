package focused_fixture_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFocused_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Focused_fixture Suite")
}
