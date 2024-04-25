package client_errors

import (
	"fmt"
)

type InvalidRegNumError struct {
	RegNum string
}

func (e *InvalidRegNumError) Error() string {
	return fmt.Sprintf("provided regNum '%s' is not valid", e.RegNum)
}

func NewInvalidRegNumError(regNum string) error {
	return &InvalidRegNumError{RegNum: regNum}
}
