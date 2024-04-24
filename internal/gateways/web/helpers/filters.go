package helpers

import (
	"fmt"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
)

const (
	URLQueryUserName       = "userName"
	URLQueryUserSurname    = "userSurname"
	URLQueryUserPatronymic = "userPatronymic"

	URLQueryCarMinYear = "minYear"
	URLQueryCarMaxYear = "maxYear"
	URLQueryCarModel   = "model"
	URLQueryCarMark    = "mark"
	URLQueryCarRegNum  = "maxYear"
)

func ParseUserFilters(ctx *gin.Context) *models.UserFilter {
	var userFilter models.UserFilter

	userFilter.Name = ctx.Param(URLQueryUserName)
	userFilter.Surname = ctx.Param(URLQueryUserSurname)
	userFilter.Patronymic = ctx.Param(URLQueryUserPatronymic)

	return &userFilter
}

func ParseCarFilters(ctx *gin.Context) *models.CarFilter {
	var carFilter models.CarFilter

	carFilter.Model = ctx.Param(URLQueryCarModel)
	carFilter.Mark = ctx.Param(URLQueryCarMark)
	carFilter.RegNum = ctx.Param(URLQueryCarRegNum)
	carFilter.Model = ctx.Param(URLQueryCarModel)

	minYear := ctx.Param(URLQueryCarMinYear)
	if minYear != "" {
		y, err := strconv.Atoi(minYear)
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("failed to parse '%s'", minYear))
		} else {
			carFilter.MinYear = int32(y)
		}
	}

	maxYear := ctx.Param(URLQueryCarMaxYear)
	if maxYear != "" {
		y, err := strconv.Atoi(maxYear)
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("failed to parse '%s'", maxYear))
		} else {
			carFilter.MinYear = int32(y)
		}
	}

	return &carFilter
}
