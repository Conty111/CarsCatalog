package car

import (
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
	"github.com/gin-gonic/gin"
)

// Check that Controller implements an apiv1.Controller
var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller is a controller implementation for status checks
type Controller struct {
	apiv1.BaseController
}

// NewController creates new status controller instance
func NewController() *Controller {
	return &Controller{
		BaseController: apiv1.BaseController{
			RelativePath: "/car",
		},
	}
}

func (ctrl *Controller) GetRelativePath() string {
	return ctrl.RelativePath
}

// GetCarsList godoc
// @Summary Get Cars List
// @Description get list of cars with pagination
// @ID get-cars
// @Accept json
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /api/v1/car/list [get]
func (ctrl *Controller) GetCarsList(ctx *gin.Context) {
	_ = helpers.ParseCarFilters(ctx)

	//render.JSONAPIPayload(ctx, http.StatusOK, &Response{
	//	Status: http.StatusText(http.StatusOK),
	//	Build:  ctrl.buildInfo,
	//})
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/list", ctrl.GetCarsList)
}
