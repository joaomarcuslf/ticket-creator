package app_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/ticket-creator/handlers"
)

func (a *AppClient) GetTicket(c *gin.Context) {
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
		c.HTML(
			http.StatusNotFound,
			"404.tmpl.html",
			nil,
		)
	} else {
		c.HTML(
			http.StatusOK,
			"ticket.tmpl.html",
			map[string]interface{}{
				"Values": map[string]interface{}{
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
