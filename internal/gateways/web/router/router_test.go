package router_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/Conty111/CarsCatalog/internal/gateways/web/router"
)

var _ = Describe("Router", func() {
	Describe("NewRouter()", func() {
		It("should create new router", func() {
			r := NewRouter()

			Expect(r).NotTo(BeNil())
		})
	})
})
