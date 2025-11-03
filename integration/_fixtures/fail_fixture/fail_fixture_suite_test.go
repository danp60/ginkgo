package fail_fixture_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFail_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fail_fixture Suite")
}
