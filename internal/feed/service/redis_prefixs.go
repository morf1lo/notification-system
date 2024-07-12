package service

import "fmt"

const (
	articlePrefix = "article:%d" // article ID
)

func ArticlePrefix(articleID int64) string {
	return fmt.Sprintf(articlePrefix, articleID)
}
