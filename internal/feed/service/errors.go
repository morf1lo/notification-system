package service

import "errors"

var (
	errMaxArticlesReached = errors.New("maximum number of articles per week reached")
)
