package helpers

import (
	"github.com/Conty111/CarsCatalog/internal/client_errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	DefaultLimit   uint = 10
	URLQueryLimit       = "limit"
	URLQueryOffset      = "offset"
)

type PaginationParams struct {
	Limit  uint
	Offset uint
}

type PaginationData struct {
	NextPage     string        `jsonapi:"next_page"`
	PreviousPage string        `jsonapi:"previous_page"`
	LastOffset   uint          `jsonapi:"last_offset"`
	Data         []interface{} `jsonapi:"data"`
}

func ParsePagination(ctx *gin.Context) (*PaginationParams, error) {
	var pag PaginationParams

	limit := ctx.Param(URLQueryLimit)
	if limit == "" {
		pag.Limit = DefaultLimit
	} else {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			return nil, err
		}
		if lim <= 0 {
			return nil, client_errors.ErrInvalidLimitParam
		}
		pag.Limit = uint(lim)
	}

	offset := ctx.Param(URLQueryOffset)
	if offset == "" {
		pag.Offset = 0
	} else {
		offs, err := strconv.Atoi(offset)
		if err != nil {
			return nil, err
		}
		if offs <= 0 {
			return nil, client_errors.ErrInvalidLimitParam
		}
		pag.Offset = uint(offs)
	}

	return &pag, nil
}
