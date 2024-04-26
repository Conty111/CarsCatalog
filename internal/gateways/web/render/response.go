package render

import (
	"errors"
	"github.com/Conty111/CarsCatalog/internal/errs"
	"github.com/Conty111/CarsCatalog/internal/external_api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func WriteErrorResponse(ctx *gin.Context, err error) {
	var status int

	switch {
	case errors.Is(err, &external_api.ExternalAPIError{}):
		status = http.StatusServiceUnavailable
	case errors.Is(err, &errs.RegNumExistError{}),
		errors.Is(err, &errs.InvalidRegNumError{}),
		errors.Is(err, &validator.ValidationErrors{}),
		errors.Is(err, errs.ErrInvalidLimitParam):

		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	ctx.AbortWithStatusJSON(status, &ErrResponse{
		Status: http.StatusText(status),
		Error:  err.Error(),
	})
}
