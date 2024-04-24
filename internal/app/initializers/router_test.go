package initializers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Conty111/CarsCatalog/internal/app/dependencies"
	. "github.com/Conty111/CarsCatalog/internal/app/initializers"
)

var _ = Describe("Router", func() {
	Describe("InitializeRouter()", func() {
		var (
			c *dependencies.Container
		)

		BeforeEach(func() {
			c = &dependencies.Container{}
		})

		It("should initialize router", func() {
			r := InitializeRouter(c)

			Expect(r).NotTo(BeNil())
		})
	})
})
