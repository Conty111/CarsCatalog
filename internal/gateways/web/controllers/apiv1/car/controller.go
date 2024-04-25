package car

import (
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/render"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Check that Controller implements an apiv1.Controller
var (
	_ apiv1.Controller = (*Controller)(nil)
)

type Service interface {
	CreateCars(regNums []string) ([]*models.Car, error)
	GetCars(pag *helpers.PaginationParams, filters *models.CarFilter) ([]models.Car, int64, error)
	GetCarByID(id uuid.UUID) (*models.Car, error)
	UpdateCarByID(id uuid.UUID, upd *helpers.CarUpdates) error
	DeleteCarByID(id uuid.UUID) error
}

// Controller is a controller implementation for status checks
type Controller struct {
	Service Service
	apiv1.BaseController
}

// NewController creates new status controller instance
func NewController(svc Service) *Controller {
	return &Controller{
		Service: svc,
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
	filters := helpers.ParseCarFilters(ctx)
	pag := helpers.ParsePagination(ctx)

	data, err := ctrl.Service.GetCars(pag, filters)
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	//var pagData *helpers.PaginationData[*car.CarInfo]
	//pagData.LastOffset = lastOffset
	//pagData.Data = make([]*car.CarInfo, len(cars))
	//
	//for i := range cars {
	//	pagData.Data[i] = &car.CarInfo{
	//		ID:     cars[i].ID.String(),
	//		Model:  cars[i].Model,
	//		Mark:   cars[i].Mark,
	//		RegNum: cars[i].RegNum,
	//		Year:   int(cars[i].Year),
	//		Owner: user.UserInfo{
	//			ID:         cars[i].OwnerID.String(),
	//			Name:       cars[i].Owner.Name,
	//			Surname:    cars[i].Owner.Surname,
	//			Patronymic: *cars[i].Owner.Patronymic,
	//		},
	//	}
	//}

	render.JSONAPIPayload(ctx, http.StatusOK, data)
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/list", ctrl.GetCarsList)
}
