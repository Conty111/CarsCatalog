package apiv1

import "github.com/gin-gonic/gin"

const BasePath = "/api/v1"

// Controller is an interface for HTTP controllers
type Controller interface {
	DefineRoutes(gin.IRouter)
	GetRelativePath() string
}

// BaseController should implement general methods for all controllers
type BaseController struct {
	RelativePath string
}
