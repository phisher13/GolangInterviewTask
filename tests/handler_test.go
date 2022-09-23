package main

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"
	"urls/pkg/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostMethod(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/v1", handlers.PostHandler())

	url := handlers.URL{
		Url: "http://test.com",
	}

	jsonValue, _ := json.Marshal(url)
    req, _ := http.NewRequest("POST", "/api/v1", bytes.NewBufferString(string(jsonValue)))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
	
	t.Logf("Calling post method. Expected 201 status code")

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestErrorPostMethod(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/v1", handlers.PostHandler())

	url := handlers.URL{
		Url: "test.com",
	}

	jsonValue, _ := json.Marshal(url)
    req, _ := http.NewRequest("POST", "/api/v1", bytes.NewBufferString(string(jsonValue)))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

	t.Logf("Calling post method with invalid url. Expected error 400")
	

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
