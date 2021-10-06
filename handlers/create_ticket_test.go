package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateTicket01(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.LoadHTMLGlob("../templates/*.tmpl.html")
	r.Static("/static", "../static")

	r.POST("/create-ticket", CreateTicket)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create-ticket", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCreateTicket02(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.LoadHTMLGlob("../templates/*.tmpl.html")
	r.Static("/static", "../static")

	r.POST("/create-ticket", CreateTicket)

	reader := strings.NewReader("title=test&description=test")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create-ticket", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusMovedPermanently, w.Code)
}
