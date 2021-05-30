module articleswebapp.com/main

go 1.16

require (
	articleswebapp.com/controllers/article_controller v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
)

replace articleswebapp.com/controllers/article_controller => ./controllers/article

replace articleswebapp.com/models/article_model => ./models/article
