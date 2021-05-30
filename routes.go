package main

import "articleswebapp.com/controllers/article_controller"

func initRoutes() {
	router.GET("/", article_controller.ShowIndexPage)
	router.GET("/article/view/:article_id", article_controller.GetArticle)
}
