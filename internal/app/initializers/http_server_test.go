package initializers_test

import (
	"github.com/Conty111/CarsCatalog/internal/configs"
	"github.com/gin-gonic/gin"

	"github.com/Conty111/CarsCatalog/internal/gateways/web/router"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/Conty111/CarsCatalog/internal/app/initializers"
)

var _ = Describe("HttpServer", func() {
	Describe("InitializeHTTPServer()", func() {
		var (
			r   *gin.Engine
			cfg *configs.HTTPServerConfig
		)

		BeforeEach(func() {
			r = router.NewRouter()
			cfg = configs.GetConfig().HTTPServer
		})

		It("should initialize HTTP server", func() {
			srv, err := InitializeHTTPServer(cfg, r)

			Expect(srv).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
