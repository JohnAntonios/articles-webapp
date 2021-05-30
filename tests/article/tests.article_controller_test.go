package article_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"articleswebapp.com/tests"
)

// Test a GET request to the home page.
// Should return the home page with the HTTP status code 200 for an unauthenticated user.

func TestShowIndexPageUnauthenticated(t *testing.T) {
	router := tests.GetRouter(true)

	// router.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	tests.TestHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)

		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
