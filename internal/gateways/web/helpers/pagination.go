package helpers

import (
	"github.com/Conty111/CarsCatalog/internal/client_errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
)

const (
	DefaultLimit   uint = 10
	DefaultOffset  uint = 0
	URLQueryLimit       = "limit"
	URLQueryOffset      = "offset"
)

type PaginationParams struct {
	Limit  uint
	Offset uint
}

type PaginationData[T any] struct {
	NextPage     string `jsonapi:"next_page"`
	PreviousPage string `jsonapi:"previous_page"`
	LastOffset   int64  `jsonapi:"last_offset"`
	Data         []T    `jsonapi:"data"`
}

func ParsePagination(ctx *gin.Context) *PaginationParams {
	var pag PaginationParams

	limit := ctx.Param(URLQueryLimit)
	if limit == "" {
		pag.Limit = DefaultLimit
	} else {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			log.Error().Err(err).Msg("error while to parsing pagination")
			pag.Limit = DefaultLimit
		} else if lim <= 0 {
			log.Error().Err(client_errors.ErrInvalidLimitParam).Msg("error while to parsing pagination")
			pag.Limit = DefaultLimit
		} else {
			pag.Limit = uint(lim)
		}
	}

	offset := ctx.Param(URLQueryOffset)
	if offset == "" {
		pag.Offset = 0
	} else {
		offs, err := strconv.Atoi(offset)
		if err != nil {
			log.Error().Err(err).Msg("error while to parsing pagination")
			pag.Offset = DefaultOffset
		} else if offs < 0 {
			log.Error().Err(client_errors.ErrInvalidLimitParam).Msg("error while to parsing pagination")
			pag.Offset = DefaultOffset
		} else {
			pag.Offset = uint(offs)
		}
	}

	return &pag
}
