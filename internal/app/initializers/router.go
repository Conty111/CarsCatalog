package initializers

import (
	"github.com/Conty111/CarsCatalog/internal/app/dependencies"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
	apiv1Cars "github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/car"
	apiv1Status "github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/status"
	apiv1Swagger "github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/swagger"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter initializes new gin router
func InitializeRouter(container *dependencies.Container) *gin.Engine {
	r := router.NewRouter()

	ctrls := buildControllers(container)

	v1 := r.Group(apiv1.BasePath)
	for i := range ctrls {
		controllerGroup := v1.Group(ctrls[i].GetRelativePath())
		ctrls[i].DefineRoutes(controllerGroup)
	}

	return r
}

func buildControllers(container *dependencies.Container) []apiv1.Controller {
	return []apiv1.Controller{
		apiv1Status.NewController(container.BuildInfo),
		apiv1Swagger.NewController(),
		apiv1Cars.NewController(container.CarService),
	}
}
