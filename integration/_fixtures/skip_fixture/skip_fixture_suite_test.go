package skip_fixture_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSkip_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Skip_fixture Suite")
}
