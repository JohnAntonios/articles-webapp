module articleswebapp.com/tests/article_test

go 1.16

replace articleswebapp.com/models/article_model => ../../models/article

replace articleswebapp.com/controllers/article_controller => ../../controllers/article

replace articleswebapp.com/tests => ../.

require (
	articleswebapp.com/models/article_model v0.0.0-00010101000000-000000000000
	articleswebapp.com/tests v0.0.0-00010101000000-000000000000
)
