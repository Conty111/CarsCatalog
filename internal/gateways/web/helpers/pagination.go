package helpers

import (
	"github.com/Conty111/CarsCatalog/internal/errs"
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

// PaginationResponse contains pagination information and data.
// swagger:model PaginationResponse
type PaginationResponse struct {
	Data           []interface{}  `json:"data"`
	PaginationMeta PaginationMeta `json:"pagination"`
} // @name PaginationResponse

// PaginationMeta contains last offset of items, nextPage (if exist) and previousPage (if exist).
type PaginationMeta struct {
	LastOffset   int64  `json:"lastOffset"`
	NextPage     string `json:"nextPage,omitempty"`
	PreviousPage string `json:"previousPage,omitempty"`
} // @name PaginationMeta

// ParsePagination gets pagination params from URL query
func ParsePagination(ctx *gin.Context) *PaginationParams {
	var pag PaginationParams

	limit := ctx.Query(URLQueryLimit)
	if limit == "" {
		pag.Limit = DefaultLimit
	} else {
		lim, err := strconv.Atoi(limit)
		if err != nil {
			log.Error().Err(err).Msg("error while to parsing pagination")
			pag.Limit = DefaultLimit
		} else if lim <= 0 {
			log.Error().Err(errs.ErrInvalidLimitParam).Msg("error while to parsing pagination")
			pag.Limit = DefaultLimit
		} else {
			pag.Limit = uint(lim)
		}
	}

	offset := ctx.Query(URLQueryOffset)
	if offset == "" {
		pag.Offset = 0
	} else {
		offs, err := strconv.Atoi(offset)
		if err != nil {
			log.Error().Err(err).Msg("error while to parsing pagination")
			pag.Offset = DefaultOffset
		} else if offs < 0 {
			log.Error().Err(errs.ErrInvalidLimitParam).Msg("error while to parsing pagination")
			pag.Offset = DefaultOffset
		} else {
			pag.Offset = uint(offs)
		}
	}

	return &pag
}
