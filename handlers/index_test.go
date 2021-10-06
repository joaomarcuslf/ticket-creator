package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func init() {
}

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.LoadHTMLGlob("../templates/*.tmpl.html")
	r.Static("/static", "../static")

	r.GET("/", Index)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
