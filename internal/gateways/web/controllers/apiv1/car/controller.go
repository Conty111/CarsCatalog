package car

import (
	"fmt"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/render"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/serializers"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Check that Controller implements an apiv1.Controller
var (
	_ apiv1.Controller = (*Controller)(nil)
)

type Service interface {
	CreateCars(regNums []string) error
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

	log.Debug().
		Str("endpoint", "GetCarsList").
		Any("pagination", pag).
		Any("filters", filters).
		Msg("parsed filter and pagination")

	carsData, lastOffset, err := ctrl.Service.GetCars(pag, filters)
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}

	log.Debug().
		Str("endpoint", "GetCarsList").
		Msg("got list of cars")

	var pagData helpers.PaginationResponse

	pagData.Data = make([]interface{}, len(carsData))
	for i := range carsData {
		pagData.Data[i] = serializers.SerializeCarInfo(&carsData[i])
	}

	pagData.PaginationMeta.LastOffset = lastOffset

	nextPage := pag.Offset + pag.Limit
	if int64(nextPage) < lastOffset {
		pagData.PaginationMeta.NextPage = fmt.Sprintf(
			"%s/list?offset=%dlimit=%d",
			ctrl.GetRelativePath(),
			nextPage,
			pag.Limit,
		)
	}

	if pag.Offset > pag.Limit {
		prevPage := pag.Offset - pag.Limit
		pagData.PaginationMeta.PreviousPage = fmt.Sprintf(
			"%s/list?offset=%d&limit=%d",
			ctrl.GetRelativePath(),
			prevPage,
			pag.Limit,
		)
	} else if pag.Offset > 0 {
		pagData.PaginationMeta.PreviousPage = fmt.Sprintf(
			"%s/list?offset=%d&limit=%d",
			ctrl.GetRelativePath(),
			0,
			pag.Limit,
		)
	}

	log.Debug().
		Str("endpoint", "GetCarsList").
		Msg("calculated pagination")

	ctx.JSON(http.StatusOK, &pagData)
}

// GetCar godoc
// @Summary Get Car
// @Description get car by id
// @ID get-car
// @Accept json
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /api/v1/car/:carID [get]
func (ctrl *Controller) GetCar(ctx *gin.Context) {
	carID, err := uuid.Parse(ctx.Param("carID"))
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	carModel, err := ctrl.Service.GetCarByID(carID)
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}

	body := serializers.SerializeCarInfo(carModel)

	ctx.JSON(http.StatusOK, body)
}

func (ctrl *Controller) CreateCars(ctx *gin.Context) {
	var body struct {
		RegNums []string `json:"regNums"`
	}
	if err := ctx.ShouldBind(&body); err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	err := ctrl.Service.CreateCars(body.RegNums)
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, &MsgResponse{
		Message: "cars successfully created",
		Status:  "OK",
	})
}

func (ctrl *Controller) DeleteCar(ctx *gin.Context) {
	carID, err := uuid.Parse(ctx.Param("carID"))
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	err = ctrl.Service.DeleteCarByID(carID)
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, &MsgResponse{
		Message: "car successfully deleted",
		Status:  "OK",
	})
}

// UpdateCar godoc
// @Summary Update Car
// @Description update car by id
// @ID update-car
// @Accept json
// @Produce json
// @Success 200 {object} MsgResponse
// @Router /api/v1/car/:carID [patch]
func (ctrl *Controller) UpdateCar(ctx *gin.Context) {
	carID, err := uuid.Parse(ctx.Param("carID"))
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}

	var upd helpers.CarUpdates
	if err = ctx.Bind(&upd); err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}

	err = ctrl.Service.UpdateCarByID(carID, &upd)
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, &MsgResponse{
		Message: "car successfully updated",
		Status:  "OK",
	})
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/list", ctrl.GetCarsList)
	r.GET("/:carID", ctrl.GetCar)
	r.POST("", ctrl.CreateCars)
	r.DELETE("/:carID", ctrl.DeleteCar)
	r.PATCH("/:carID", ctrl.UpdateCar)
}
