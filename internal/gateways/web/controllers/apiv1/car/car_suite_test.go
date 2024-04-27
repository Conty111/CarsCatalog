package car_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Car Suite")
}
