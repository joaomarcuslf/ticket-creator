package rest_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *RestClient) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
