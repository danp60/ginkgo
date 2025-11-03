package testingtproxy_test

import (
	"testing"

	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestingtproxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testingtproxy Suite")
}
