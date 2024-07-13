package service

import "fmt"

const (
	articlePrefix = "article:%d" // article ID
	articlesCooldownPrefix = "articles-cooldown"
)

func ArticlePrefix(articleID int64) string {
	return fmt.Sprintf(articlePrefix, articleID)
}
