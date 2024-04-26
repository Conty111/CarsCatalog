package render

import (
	"errors"
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

	switch err {
	default:
		status = http.StatusInternalServerError
	}

	if errors.As(err, &validator.ValidationErrors{}) {
		status = http.StatusBadRequest
	}
	ctx.AbortWithStatusJSON(status, &ErrResponse{
		Status: http.StatusText(status),
		Error:  err.Error(),
	})
}
