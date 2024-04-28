package helpers

import (
	"fmt"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"regexp"
	"strconv"
)

const (
	URLQueryUserName       = "userName"
	URLQueryUserSurname    = "userSurname"
	URLQueryUserPatronymic = "userPatronymic"

	CarRegNumPattern = "^[ABEKMHOPCTYXАВЕКМНОРСТУХ]{1}\\d{3}[ABEKMHOPCTYXАВЕКМНОРСТУХ]{2}\\d{2,3}$"

	URLQueryCarMinYear = "minYear"
	URLQueryCarMaxYear = "maxYear"
	URLQueryCarModel   = "model"
	URLQueryCarMark    = "mark"
	URLQueryCarRegNum  = "maxYear"
)

func ParseUserFilters(ctx *gin.Context) *models.UserFilter {
	var userFilter models.UserFilter

	userFilter.Name = ctx.Query(URLQueryUserName)
	userFilter.Surname = ctx.Query(URLQueryUserSurname)
	userFilter.Patronymic = ctx.Query(URLQueryUserPatronymic)

	return &userFilter
}

func ParseCarFilters(ctx *gin.Context) *models.CarFilter {
	var carFilter models.CarFilter

	carFilter.Model = ctx.Query(URLQueryCarModel)
	carFilter.Mark = ctx.Query(URLQueryCarMark)
	carFilter.Model = ctx.Query(URLQueryCarModel)

	re := regexp.MustCompile(CarRegNumPattern)
	regNum := ctx.Query(URLQueryCarRegNum)
	if !re.MatchString(regNum) {
		log.Error().Msg("invalid reg num in params")
	} else {
		carFilter.RegNum = regNum
	}

	minYear := ctx.Query(URLQueryCarMinYear)
	if minYear != "" {
		y, err := strconv.Atoi(minYear)
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("failed to parse '%s'", minYear))
		} else {
			carFilter.MinYear = int32(y)
		}
	}

	maxYear := ctx.Query(URLQueryCarMaxYear)
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
