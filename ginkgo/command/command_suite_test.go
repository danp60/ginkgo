package command_test

import (
	"testing"

	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Command Suite")
}
