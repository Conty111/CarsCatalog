package initializers

import (
	"fmt"
	"github.com/Conty111/CarsCatalog/internal/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitializeHTTPServer create new http.Server instance
func InitializeHTTPServer(cfg *configs.HTTPServerConfig, router *gin.Engine) (*http.Server, error) {
	// create http server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: router,
	}
	return srv, nil
}
