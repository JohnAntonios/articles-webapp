package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"articleswebapp.com/models/article_model"
	"github.com/gin-gonic/gin"
)

var tmpArticleList []article_model.Article

// This function is used for setup before executing the test functions.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing.
func GetRouter(withTemplates bool) *gin.Engine {
	engine := gin.Default()

	if withTemplates {
		engine.LoadHTMLGlob("../../templates/*")
	}

	return engine
}

// Helper function to process a request and test its response.
func TestHTTPResponse(t *testing.T, e *gin.Engine, req *http.Request, controller func(w *httptest.ResponseRecorder) bool) {
	responseRecorder := httptest.NewRecorder()

	// Create the service and process the incoming request.
	e.ServeHTTP(responseRecorder, req)

	if !controller(responseRecorder) {
		t.Fail()
	}
}

// This function stores the main list into a temporary one for testing.
func SaveLists() {
	tmpArticleList = article_model.GetAllArticles()
}
