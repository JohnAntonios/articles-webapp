package article_model

import (
	"fmt"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 Body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 Body"},
}

func GetAllArticles() []Article {
	return articleList
}

func GetArticleById(articleId int) (*Article, error) {
	for _, article := range articleList {
		if articleId == article.ID {
			return &article, nil
		}
	}
	return nil, fmt.Errorf("Article not found")
}
