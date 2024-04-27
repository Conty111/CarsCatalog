package car_test

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

func BeforeCarSuite() {
	if envy.Get("ENABLE_JSON_LOGS", "false") == "false" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	gin.SetMode(gin.ReleaseMode)
}

func TestCar(t *testing.T) {
	RegisterFailHandler(Fail)
	BeforeSuite(BeforeCarSuite)
	RunSpecs(t, "Car Suite")
}
