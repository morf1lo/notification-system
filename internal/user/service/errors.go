package service

import "errors"

var (
	errUserIsAlreadyExists = errors.New("user with this email is already exists")
	errMaxSubscribersCountReached = errors.New("maximum count of subscribers reached")
)
