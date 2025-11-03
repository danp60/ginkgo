package tags_test

import (
	. "github.com/danp60/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTagsTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TagsTests Suite")
}
