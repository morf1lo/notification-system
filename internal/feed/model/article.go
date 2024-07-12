package model

type Article struct {
	Title  string `json:"title" binding:"required,min=10"`
	Body   string `json:"body" binding:"required,min=120"`
	Author string `json:"author" binding:"required"`
}
