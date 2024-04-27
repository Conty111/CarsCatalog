package car_test

import (
	"encoding/json"
	. "github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/car"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/Conty111/CarsCatalog/internal/services"
	"github.com/Conty111/CarsCatalog/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Car Controller", func() {
	var (
		t            GinkgoTInterface
		w            *httptest.ResponseRecorder
		ctx          *gin.Context
		carManager   *mocks.CarManager
		userProvider *mocks.UserProvider
		extApiClient *mocks.ExternalAPIClient
		service      Service
		carCtrl      *Controller

		owner             *models.User
		ladaVesta         *models.Car
		mercedes          *models.Car
		volgaWithoutOwner *models.Car
	)
	BeforeEach(func() {
		w = httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)

		t = GinkgoT()

		carManager = mocks.NewCarManager(t)
		userProvider = mocks.NewUserProvider(t)
		extApiClient = mocks.NewExternalAPIClient(t)

		service = services.NewCarService(carManager, extApiClient, userProvider)

		carCtrl = NewController(service)

		patr := "Бенджамин"
		owner = &models.User{
			Name:       "Питер",
			Surname:    "Паркер",
			Patronymic: &patr,
		}
		owner.ID = uuid.New()

		ladaVesta = &models.Car{
			RegNum:  "A123AA12",
			Mark:    "Lada",
			Model:   "Vesta",
			Year:    2010,
			OwnerID: &owner.ID,
			Owner:   owner,
		}
		ladaVesta.ID = uuid.New()

		mercedes = &models.Car{
			RegNum:  "A777AA77",
			Model:   "Mercedes",
			Mark:    "Benz",
			Year:    2020,
			OwnerID: &owner.ID,
			Owner:   owner,
		}
		owner.Cars = []*models.Car{mercedes, ladaVesta}

		volgaWithoutOwner = &models.Car{
			RegNum: "A111BC11",
			Mark:   "Volga",
			Model:  "Best",
			Year:   2003,
		}
		volgaWithoutOwner.ID = uuid.New()
	})

	It("controller should not be nil", func() {
		Expect(carCtrl).NotTo(BeNil())
	})

	Describe("GetCarsList()", func() {
		Context("request without filters", func() {
			It("should return list of cars with pagination meta", func() {
				var offs, limit int = 0, 10
				allCars := []models.Car{*mercedes, *ladaVesta, *volgaWithoutOwner}

				carManager.Mock.
					On("GetCars", offs, limit, mock.AnythingOfType("*models.CarFilter")).
					Once().Return(allCars, nil)
				carManager.Mock.
					On("GetLastOffset", mock.AnythingOfType("*models.CarFilter")).
					Once().Return(int64(len(allCars)), nil)

				ctx.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/car/list", nil)
				carCtrl.GetCarsList(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

				responseBody, err := io.ReadAll(w.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				var pagData helpers.PaginationResponse
				err = json.Unmarshal(responseBody, &pagData)
				Expect(err).NotTo(HaveOccurred())

				Expect(len(pagData.Data)).To(Equal(len(allCars)))
				Expect(pagData.PaginationMeta.LastOffset).To(Equal(int64(len(allCars))))
			})
		})
		Context("request with URL query params", func() {
			It("should return status OK and pagination meta", func() {
				var offs, limit int = 0, 10
				allCars := []models.Car{*mercedes}

				carManager.Mock.
					On("GetCars", offs, limit, mock.AnythingOfType("*models.CarFilter")).
					Once().Return(allCars, nil)
				carManager.Mock.
					On("GetLastOffset", mock.AnythingOfType("*models.CarFilter")).
					Once().Return(int64(len(allCars)), nil)

				params := url.Values{}
				params.Set("limit", strconv.Itoa(limit))
				params.Set("offset", strconv.Itoa(offs))
				params.Set("model", mercedes.Model)
				params.Set("mark", mercedes.Mark)
				params.Set("regNum", mercedes.RegNum)
				params.Set("minYear", strconv.Itoa(int(mercedes.Year-1)))
				params.Set("maxYear", strconv.Itoa(int(mercedes.Year+1)))

				req, _ := http.NewRequest(http.MethodGet, "/api/v1/car/list"+"?"+params.Encode(), nil)

				ctx.Request = req

				carCtrl.GetCarsList(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

				responseBody, err := io.ReadAll(w.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				var pagData helpers.PaginationResponse
				err = json.Unmarshal(responseBody, &pagData)
				Expect(err).NotTo(HaveOccurred())

				Expect(len(pagData.Data)).To(Equal(len(allCars)))
				Expect(pagData.PaginationMeta.LastOffset).To(Equal(int64(len(allCars))))
			})
		})
	})
})
