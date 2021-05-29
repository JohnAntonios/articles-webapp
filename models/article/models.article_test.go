package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	testArticleList := getAllArticles()

	if len(testArticleList) != len(articleList) {
		t.Fail()
	}

	for index, article := range testArticleList {
		currArticle := articleList[index]
		if article.ID != currArticle.ID ||
			article.Title != currArticle.Title ||
			article.Content != currArticle.Content {
			t.Fail()
			break
		}
	}

}
