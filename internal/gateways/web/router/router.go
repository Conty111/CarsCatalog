package router

import (
	"github.com/gin-gonic/gin"

	ginmw "github.com/Conty111/CarsCatalog/internal/gateways/web/middleware"
)

// NewRouter create new gin router instance
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(ginmw.Recovery(), ginmw.Logger())

	return r
}
