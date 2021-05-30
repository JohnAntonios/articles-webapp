package article_controller

import (
	"net/http"
	"strconv"

	"articleswebapp.com/models/article_model"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := article_model.GetAllArticles()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func GetArticle(c *gin.Context) {
	if articleId, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := article_model.GetArticleById(articleId); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
