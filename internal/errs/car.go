package errs

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrInvalidBody = errors.New("invalid request body provided")
)

type CarNotFoundError struct {
	CarID uuid.UUID
}

func (e *CarNotFoundError) Error() string {
	return fmt.Sprintf("car with provided car ID '%s' not found", e.CarID.String())
}
func NewCarNotFoundError(carID uuid.UUID) error {
	return &CarNotFoundError{CarID: carID}
}

type InvalidRegNumError struct {
	RegNum string
}

func (e *InvalidRegNumError) Error() string {
	return fmt.Sprintf("provided regNum '%s' is not valid", e.RegNum)
}
func NewInvalidRegNumError(regNum string) error {
	return &InvalidRegNumError{RegNum: regNum}
}

type RegNumExistError struct {
	RegNum string
}

func (e *RegNumExistError) Error() string {
	return fmt.Sprintf("provided regNum '%s' is already exist", e.RegNum)
}
func NewRegNumExistError(regNum string) error {
	return &RegNumExistError{RegNum: regNum}
}
