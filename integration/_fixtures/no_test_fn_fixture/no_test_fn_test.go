package no_test_fn_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/danp60/ginkgo/v2/integration/_fixtures/no_test_fn_fixture"
	. "github.com/onsi/gomega"
)

var _ = Describe("NoTestFn", func() {
	It("should proxy strings", func() {
		Ω(StringIdentity("foo")).Should(Equal("foo"))
	})
})
