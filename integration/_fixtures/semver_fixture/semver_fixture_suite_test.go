package semver_fixture_test

import (
	"testing"

	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSemverFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SemverFixture Suite")
}
