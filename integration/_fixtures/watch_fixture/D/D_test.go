package D_test

import (
	. "github.com/danp60/ginkgo/v2/integration/_fixtures/watch_fixture/D"

	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
