module articleswebapp.com/tests

go 1.16

replace articleswebapp.com/models/article_model => ../models/article

require (
	articleswebapp.com/models/article_model v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
)
