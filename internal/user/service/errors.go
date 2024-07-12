package service

import "errors"

var (
	errUserIsAlreadyExists = errors.New("user with this email is already exists")
)
