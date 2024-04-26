package status

import (
	"github.com/Conty111/CarsCatalog/internal/app/build"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
	"github.com/gin-gonic/gin"

	"net/http"
)

// Check that Controller implements an apiv1.Controller
var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller is a controller implementation for status checks
type Controller struct {
	apiv1.BaseController
	buildInfo *build.Info
}

// NewController creates new status controller instance
func NewController(bi *build.Info) *Controller {
	return &Controller{
		buildInfo: bi,
		BaseController: apiv1.BaseController{
			RelativePath: "/status",
		},
	}
}

func (ctrl *Controller) GetRelativePath() string {
	return ctrl.RelativePath
}

// GetStatus godoc
// @Summary Get Application Status
// @Description get status
// @ID get-status
// @Accept json
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /api/v1/status [get]
func (ctrl *Controller) GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Response{
		Status: http.StatusText(http.StatusOK),
		Build:  ctrl.buildInfo,
	})
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("", ctrl.GetStatus)
}
