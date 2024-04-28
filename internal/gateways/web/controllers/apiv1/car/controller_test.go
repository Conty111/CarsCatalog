package car_test

import (
	"bytes"
	"encoding/json"
	"github.com/Conty111/CarsCatalog/internal/errs"
	"github.com/Conty111/CarsCatalog/internal/external_api"
	. "github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/car"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/serializers"
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
	const baseURI = "/api/v1"

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
		mercedes.ID = uuid.New()
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

				ctx.Request, _ = http.NewRequest(http.MethodGet, baseURI+"/car/list", nil)
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

				req, _ := http.NewRequest(http.MethodGet, baseURI+"/car/list"+"?"+params.Encode(), nil)

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
		Context("request with invalid URL query params", func() {
			It("should return status OK with default pagination meta", func() {
				var offs, limit int = -1, 10
				allCars := []models.Car{*mercedes}

				carManager.Mock.
					On("GetCars", 0, limit, mock.AnythingOfType("*models.CarFilter")).
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

				req, _ := http.NewRequest(http.MethodGet, baseURI+"/car/list"+"?"+params.Encode(), nil)

				ctx.Request = req

				carCtrl.GetCarsList(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
	Describe("GetCar()", func() {
		Context("valid car ID provided", func() {
			It("should return car information with status OK", func() {
				carManager.Mock.
					On("GetByID", mercedes.ID).
					Once().Return(mercedes, nil)

				ctx.Request, _ = http.NewRequest(http.MethodGet, baseURI+"/car/", nil)
				ctx.AddParam("carID", mercedes.ID.String())

				carCtrl.GetCar(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

				responseBody, err := io.ReadAll(w.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var carInfo serializers.CarInfo
				err = json.Unmarshal(responseBody, &carInfo)
				Expect(err).NotTo(HaveOccurred())

				Expect(carInfo.ID).To(Equal(mercedes.ID.String()))
				Expect(carInfo.RegNum).To(Equal(mercedes.RegNum))
				Expect(carInfo.Model).To(Equal(mercedes.Model))
				Expect(carInfo.Mark).To(Equal(mercedes.Mark))
				Expect(carInfo.Year).To(Equal(int(mercedes.Year)))
			})
		})

		Context("invalid car ID provided", func() {
			It("should return status BadRequest", func() {
				invalidCarID := "invalid-car-id"

				ctx.Request, _ = http.NewRequest(http.MethodGet, baseURI+"/car/", nil)
				ctx.AddParam("carID", invalidCarID)

				carCtrl.GetCar(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
			})
		})

		Context("car not found", func() {
			It("should return status NotFound", func() {
				notFoundCarID := uuid.New()

				carManager.Mock.
					On("GetByID", notFoundCarID).
					Once().Return(nil, errs.NewCarNotFoundError(notFoundCarID))

				ctx.Request, _ = http.NewRequest(http.MethodGet, baseURI+"/car/", nil)
				ctx.AddParam("carID", notFoundCarID.String())

				carCtrl.GetCar(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
	Describe("CreateCars()", func() {
		Context("valid request body provided", func() {
			It("should return status Created with success message", func() {
				regNums := []string{volgaWithoutOwner.RegNum, ladaVesta.RegNum}
				reqBody, _ := json.Marshal(map[string][]string{"regNums": regNums})

				req, _ := http.NewRequest(http.MethodPost, baseURI+"/car", bytes.NewReader(reqBody))
				req.Header.Set("Content-Type", "application/json")
				ctx.Request = req

				extApiClient.Mock.
					On("GetCarInfo", volgaWithoutOwner.RegNum).
					Once().Return(
					&external_api.CarData{
						RegNum: volgaWithoutOwner.RegNum,
						Mark:   volgaWithoutOwner.Mark,
						Model:  volgaWithoutOwner.Model,
						Year:   int(volgaWithoutOwner.Year),
					}, nil)

				extApiClient.Mock.
					On("GetCarInfo", ladaVesta.RegNum).
					Once().Return(&external_api.CarData{
					RegNum: ladaVesta.RegNum,
					Mark:   ladaVesta.Mark,
					Model:  ladaVesta.Model,
					Year:   int(ladaVesta.Year),
					Owner: &external_api.PeopleData{
						Name:       owner.Name,
						Surname:    owner.Surname,
						Patronymic: *owner.Patronymic,
					},
				}, nil)

				userProvider.Mock.
					On("GetByFullName", owner.Name, owner.Surname, *owner.Patronymic).
					Once().Return(owner, nil)

				carManager.Mock.
					On("CreateCars", []*models.Car{
						{
							RegNum: volgaWithoutOwner.RegNum,
							Mark:   volgaWithoutOwner.Mark,
							Model:  volgaWithoutOwner.Model,
							Year:   volgaWithoutOwner.Year,
						},
						{
							RegNum:  ladaVesta.RegNum,
							Mark:    ladaVesta.Mark,
							Model:   ladaVesta.Model,
							Year:    ladaVesta.Year,
							OwnerID: &owner.ID,
							Owner:   owner,
						},
					}).
					Once().Return(nil)

				carCtrl.CreateCars(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusCreated))
			})
		})

		Context("invalid request body provided", func() {
			It("should return status BadRequest", func() {
				invalidReqBody := []byte(`{"invalidField": "value"}`)

				req, _ := http.NewRequest(http.MethodPost, baseURI+"/car", bytes.NewReader(invalidReqBody))
				req.Header.Set("Content-Type", "application/json")
				ctx.Request = req

				carCtrl.CreateCars(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
			})
		})
	})

	Describe("DeleteCar()", func() {
		Context("valid car ID provided", func() {
			It("should return status OK with success message", func() {
				ctx.Request, _ = http.NewRequest(http.MethodDelete, baseURI+"/car/"+mercedes.ID.String(), nil)
				ctx.AddParam("carID", mercedes.ID.String())

				carManager.Mock.
					On("DeleteByID", mercedes.ID).
					Once().Return(nil)

				carCtrl.DeleteCar(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

				responseBody, err := io.ReadAll(w.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var msgResponse MsgResponse
				err = json.Unmarshal(responseBody, &msgResponse)
				Expect(err).NotTo(HaveOccurred())

				Expect(msgResponse.Status).To(Equal("OK"))
			})
		})

		Context("invalid car ID provided", func() {
			It("should return status BadRequest", func() {
				invalidCarID := "invalid-car-id"
				ctx.Request, _ = http.NewRequest(http.MethodDelete, baseURI+"/car/"+invalidCarID, nil)
				ctx.AddParam("carID", invalidCarID)
				carCtrl.DeleteCar(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
			})
		})

		Context("car not found", func() {
			It("should return status NotFound", func() {
				notFoundCarID := uuid.New()
				carManager.Mock.
					On("DeleteByID", notFoundCarID).
					Once().Return(errs.NewCarNotFoundError(notFoundCarID))

				ctx.Request, _ = http.NewRequest(http.MethodDelete, baseURI+"/car/"+notFoundCarID.String(), nil)
				ctx.AddParam("carID", notFoundCarID.String())
				carCtrl.DeleteCar(ctx)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

	//Describe("UpdateCar()", func() {
	//	Context("valid car ID and update provided", func() {
	//		It("should return status OK with success message", func() {
	//			carID := mercedes.ID // Assuming 'mercedes' is the car to be updated
	//			update := CarUpdates{
	//				Model: "New Model",
	//				Mark:  "New Mark",
	//				Year:  2025,
	//			}
	//
	//			reqBody, _ := json.Marshal(update)
	//			reqURL := "/cars/" + carID.String()
	//			req, _ := http.NewRequest(http.MethodPut, reqURL, bytes.NewReader(reqBody))
	//			req.Header.Set("Content-Type", "application/json")
	//			ctx.Request = req
	//
	//			carManager.Mock.
	//				On("UpdateCarByID", carID, &update).
	//				Once().Return(nil)
	//
	//			carCtrl.UpdateCar(ctx)
	//
	//			Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
	//
	//			responseBody, err := io.ReadAll(w.Result().Body)
	//			Expect(err).NotTo(HaveOccurred())
	//
	//			var msgResponse MsgResponse
	//			err = json.Unmarshal(responseBody, &msgResponse)
	//			Expect(err).NotTo(HaveOccurred())
	//
	//			Expect(msgResponse.Message).To(Equal("car successfully updated"))
	//			Expect(msgResponse.Status).To(Equal("OK"))
	//		})
	//	})
	//
	//	Context("invalid car ID provided", func() {
	//		It("should return status BadRequest", func() {
	//			invalidCarID := "invalid-car-id"
	//			update := CarUpdates{
	//				Model: "New Model",
	//				Mark:  "New Mark",
	//				Year:  2025,
	//			}
	//
	//			reqBody, _ := json.Marshal(update)
	//			reqURL := "/cars/" + invalidCarID
	//			req, _ := http.NewRequest(http.MethodPut, reqURL, bytes.NewReader(reqBody))
	//			req.Header.Set("Content-Type", "application/json")
	//			ctx.Request = req
	//
	//			carCtrl.UpdateCar(ctx)
	//
	//			Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
	//		})
	//	})
	//
	//	Context("car not found", func() {
	//		It("should return status NotFound", func() {
	//			notFoundCarID := uuid.New()
	//			update := CarUpdates{
	//				Model: "New Model",
	//				Mark:  "New Mark",
	//				Year:  2025,
	//			}
	//
	//			reqBody, _ := json.Marshal(update)
	//			reqURL := "/cars/" + notFoundCarID.String()
	//			req, _ := http.NewRequest(http.MethodPut, reqURL, bytes.NewReader(reqBody))
	//			req.Header.Set("Content-Type", "application/json")
	//			ctx.Request = req
	//
	//			carManager.Mock.
	//				On("UpdateCarByID", notFoundCarID, &update).
	//				Once().Return(errors.New("car not found"))
	//
	//			carCtrl.UpdateCar(ctx)
	//
	//			Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
	//		})
	//	})
	//
	//	Context("invalid request body provided", func() {
	//		It("should return status BadRequest", func() {
	//			invalidReqBody := []byte(`{"invalidField": "value"}`)
	//			carID := mercedes.ID // Assuming 'mercedes' is the car to be updated
	//
	//			reqURL := "/cars/" + carID.String()
	//			req, _ := http.NewRequest(http.MethodPut, reqURL, bytes.NewReader(invalidReqBody))
	//			req.Header.Set("Content-Type", "application/json")
	//			ctx.Request = req
	//
	//			carCtrl.UpdateCar(ctx)
	//
	//			Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
	//		})
	//	})
	//})

})
