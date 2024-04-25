package repositories

import "errors"

var (
	UserNotFound = errors.New("user not found")
	CarNotFound  = errors.New("car not found")
)
