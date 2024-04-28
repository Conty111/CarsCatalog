package render

import (
	"errors"
	"github.com/Conty111/CarsCatalog/internal/errs"
	"github.com/Conty111/CarsCatalog/internal/external_api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

type ErrResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func WriteErrorResponse(ctx *gin.Context, err error) {
	var status int

	var carNotFoundError *errs.CarNotFoundError
	var externalAPIError *external_api.ExternalAPIError
	var regNumExistError *errs.RegNumExistError
	var invalidRegNumError *errs.InvalidRegNumError
	var validationErrors *validator.ValidationErrors

	switch {
	case errors.As(err, &carNotFoundError):
		status = http.StatusNotFound

	case errors.As(err, &externalAPIError):
		status = http.StatusServiceUnavailable

	case errors.As(err, &regNumExistError),
		errors.As(err, &invalidRegNumError),
		errors.As(err, &validationErrors),
		errors.As(err, &validator.ValidationErrors{}),
		errors.Is(err, errs.ErrInvalidLimitParam),
		errors.Is(err, errs.ErrInvalidOffsetParam),
		errors.Is(err, errs.ErrInvalidBody),
		uuid.IsInvalidLengthError(err):

		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	ctx.AbortWithStatusJSON(status, &ErrResponse{
		Status: http.StatusText(status),
		Error:  err.Error(),
	})
}
