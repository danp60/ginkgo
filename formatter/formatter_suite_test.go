package formatter_test

import (
	"testing"

	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFormatter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Formatter Suite")
}
