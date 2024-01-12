package gin_sample

import (
	"net/http"
	"net/http/httptest"
	"testing"
	// "github.com/gin-gonic/gin"
	// "github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
	if w.Body.String() != "Hello, Gin!" {
		t.Errorf("Home page didn't return %v", "Hello, Gin!")
	}
}

func TestHelloNameRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello/gin", nil)
	router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "Hello, gin!", w.Body.String())
}

func TestSearchRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search?q=test", nil)
	router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "Searched for: test", w.Body.String())
}
