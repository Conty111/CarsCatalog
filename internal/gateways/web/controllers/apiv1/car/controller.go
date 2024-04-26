package car

import (
	"fmt"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1"
	. "github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
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
	GetCars(pag *PaginationParams, filters *models.CarFilter) ([]models.Car, int64, error)
	GetCarByID(id uuid.UUID) (*models.Car, error)
	UpdateCarByID(id uuid.UUID, upd *CarUpdates) error
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
// @Summary Get a list of cars
// @Description Get a paginated list of cars based on filters
// @ID get-cars-list
// @Produce json
// @Param limit query int false "Limit number of items per page"
// @Param offset query int false "Offset for pagination"
// @Param model query string false "Model of the car"
// @Param mark query string false "Mark of the car"
// @Param regNum query string false "Registration number of the car"
// @Param minYear query int false "Minimum manufacturing year of the car"
// @Param maxYear query int false "Maximum manufacturing year of the car"
// @Success 200 {object} PaginationResponse "Success response"
// @Failure 400 {object} render.ErrResponse "Bad request"
// @Failure 500 {object} render.ErrResponse "Internal server error"
// @Router /car/list [get]
func (ctrl *Controller) GetCarsList(ctx *gin.Context) {
	filters := ParseCarFilters(ctx)
	pag := ParsePagination(ctx)

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

	var pagData PaginationResponse

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
// @Summary Retrieve information about a car
// @Description Retrieve information about a car using its id
// @ID get-car-info
// @Param carID query string true "ID of the car"
// @Success 200 {object} serializers.CarInfo "Success Car Info"
// @Failure 400 {object} render.ErrResponse "Bad request"
// @Failure 500 {object} render.ErrResponse "Internal server error"
// @Router /car/{carID} [get]
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

// CreateCars godoc
// @Summary Create cars
// @Description Create cars using their registration numbers
// @ID create-cars
// @Accept json
// @Produce json
// @Param regNums body []string true "Array of registration numbers of the cars"
// @Success 201 {object} MsgResponse "Success response"
// @Failure 400 {object} render.ErrResponse "Bad request"
// @Failure 500 {object} render.ErrResponse "Internal server error"
// @Router /cars [post]
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

// DeleteCar godoc
// @Summary Delete a car by ID
// @Description Delete a car using its ID
// @ID delete-car-by-id
// @Param carID path string true "ID of the car to delete"
// @Produce json
// @Success 200 {object} MsgResponse "Success response"
// @Failure 400 {object} render.ErrResponse "Bad request"
// @Failure 404 {object} render.ErrResponse "Not found"
// @Failure 500 {object} render.ErrResponse "Internal server error"
// @Router /cars/{carID} [delete]
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
// @Summary Update a car by ID
// @Description Update a car using its ID
// @ID update-car-by-id
// @Param carID path string true "ID of the car to update"
// @Accept json
// @Produce json
// @Param updateCar body CarUpdates true "Car updates"
// @Success 200 {object} MsgResponse "Success response"
// @Failure 400 {object} render.ErrResponse "Bad request"
// @Failure 404 {object} render.ErrResponse "Not found"
// @Failure 500 {object} render.ErrResponse "Internal server error"
// @Router /cars/{carID} [put]
func (ctrl *Controller) UpdateCar(ctx *gin.Context) {
	carID, err := uuid.Parse(ctx.Param("carID"))
	if err != nil {
		render.WriteErrorResponse(ctx, err)
		return
	}

	var upd CarUpdates
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
