package article_test

import (
	"testing"

	"articleswebapp.com/models/article_model"
)

func TestGetAllArticles(t *testing.T) {
	articleList := article_model.GetAllArticles()

	if len(articleList) == 0 {
		t.Fail()
	}
}
