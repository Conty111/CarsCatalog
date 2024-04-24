package build_test

import (
	. "github.com/Conty111/CarsCatalog/internal/app/build"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info", func() {
	Describe("NewInfo()", func() {
		It("should create new info object", func() {
			info := NewInfo()

			Expect(info).NotTo(BeNil())
		})
	})
})
