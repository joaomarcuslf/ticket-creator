package app_client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetTicket01(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.LoadHTMLGlob("../../templates/*.tmpl.html")
	r.Static("/static", "../../static")

	AppClient := NewAppClient("80")

	r.GET("/ticket/aaaaaa", AppClient.GetTicket)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ticket/aaaaaa", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetTicket02(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.LoadHTMLGlob("../../templates/*.tmpl.html")
	r.Static("/static", "../../static")

	AppClient := NewAppClient("80")

	r.GET("/ticket/:encodedUrl", AppClient.GetTicket)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ticket/VGVzdGUtLXwtLXRlc3RlLS18LS0yMDIxLTEwLTA2", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
