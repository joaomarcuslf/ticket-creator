package rest_client

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

	RestClient := NewRestClient("80")

	r.POST("/ticket", RestClient.CreateTicket)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ticket", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateTicket02(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	RestClient := NewRestClient("80")

	r.POST("/ticket", RestClient.CreateTicket)

	bodyReader := strings.NewReader(`{"title": "", "description": ""}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ticket", bodyReader)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCreateTicket03(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	RestClient := NewRestClient("80")

	r.POST("/ticket", RestClient.CreateTicket)

	bodyReader := strings.NewReader(`{"title": "Test", "description": "test"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ticket", bodyReader)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
