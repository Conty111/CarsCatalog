package render

import (
	"errors"
	"github.com/Conty111/CarsCatalog/internal/client_errors"
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
	//case errors.Is(err, client_errors.):

	case errors.As(err, &client_errors.RegNumExistError{}),
		errors.As(err, &client_errors.InvalidRegNumError{}),
		errors.As(err, &validator.ValidationErrors{}),
		errors.Is(err, client_errors.ErrInvalidLimitParam):

		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	ctx.AbortWithStatusJSON(status, &ErrResponse{
		Status: http.StatusText(status),
		Error:  err.Error(),
	})
}
