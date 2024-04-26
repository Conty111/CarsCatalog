package client_errors

import "errors"

var (
	UserNotFound = errors.New("user not found")
)
