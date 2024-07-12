package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Subscriber struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	DateAdded time.Time `json:"dateAdded"`
}

func (s *Subscriber) Validate() error {
	return validator.New().Struct(s)
}
