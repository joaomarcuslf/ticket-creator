package app_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *AppClient) Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		nil,
	)
}
