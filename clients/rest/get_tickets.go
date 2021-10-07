package rest_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/ticket-creator/handlers"
)

func (a *RestClient) GetTicket(c *gin.Context) {
	encodedUrl := c.Param("encodedUrl")

	scheme := "http"

	if c.Request.TLS != nil {
		scheme = "https"
	}

	ticket, err := handlers.GetTicket(
		encodedUrl,
		scheme,
		c.Request.Host,
		c.Request.URL.Path,
	)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"values": map[string]interface{}{
					"title":            ticket.Title,
					"description":      ticket.Description,
					"description_safe": ticket.SafeDescription,
					"date":             ticket.Date,
					"short_url":        ticket.ShortUrl,
				},
			},
		)
	}
}
