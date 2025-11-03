package A_test

import (
	. "github.com/danp60/ginkgo/v2/integration/_fixtures/watch_fixture/A"

	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
