package model

import "time"

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required,min=10"`
	Body      string    `json:"body" binding:"required,min=120"`
	Author    string    `json:"author" binding:"required"`
	DateAdded time.Time `json:"dateAdded"`
}
