package logic

import (
	"time"

	"github.com/BKR-dev/AtollHex/util"
)

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func GetArticlesFromPage(page int) []util.Article {
	pageSize := 20
	articles := util.GetArticles()
	start := (page - 1) * pageSize
	end := start + pageSize

	if end > len(articles) {
		end = len(articles)
	}

	return articles[start:end]
}
