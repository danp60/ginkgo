package eventually_failing_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEventuallyFailing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EventuallyFailing Suite")
}
