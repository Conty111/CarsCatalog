package client_errors

import "errors"

var (
	ErrInvalidLimitParam  = errors.New("invalid limit param")
	ErrInvalidOffsetParam = errors.New("invalid offset param")
)
