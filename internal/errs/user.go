package errs

import "errors"

var (
	UserNotFound = errors.New("user not found")
)
