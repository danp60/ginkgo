package passing_ginkgo_tests_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPassing_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Passing_ginkgo_tests Suite")
}
