package coverage_fixture_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CoverageFixture Suite")
}
