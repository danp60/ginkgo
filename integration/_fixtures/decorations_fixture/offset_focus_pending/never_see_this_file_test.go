package decorations_fixture_test

import . "github.com/danp60/ginkgo/v2"

func OffsetIt() {
	It("is offset", Offset(1), func() {
	})
}
