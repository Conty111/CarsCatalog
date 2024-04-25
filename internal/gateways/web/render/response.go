package render

import (
	"errors"
	//. "github.com/Conty111/CarsCatalog/internal/client_errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func WriteErrorResponse(ctx *gin.Context, err error) {
	if errors.As(err, &validator.ValidationErrors{}) {
		BadRequest(ctx, err.Error())
	}

	switch err {
	default:
		InternalServerError(ctx, err.Error())
	}
}
