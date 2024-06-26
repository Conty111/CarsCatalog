package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	//nolint: golint //reason: blank import because of swagger docs init
	_ "github.com/Conty111/CarsCatalog/docs/api/web"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller implements controller for swagger
type Controller struct {
	apiv1.BaseController
}

// NewController create new instance for swagger controller
func NewController() *Controller {
	return &Controller{
		BaseController: apiv1.BaseController{
			RelativePath: "/swagger",
		},
	}
}

func (ctrl *Controller) GetRelativePath() string {
	return ctrl.RelativePath
}

// DefineRoutes adds swagger controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
