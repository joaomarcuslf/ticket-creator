package rest_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/ticket-creator/usecases"
)

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (a *RestClient) CreateTicket(c *gin.Context) {
	var requestData CreateTicketRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errors := usecases.ValidateForm(
		requestData.Title,
		requestData.Description,
	)

	if errors.Ok {
		encoded := usecases.EncodeTicketData(requestData.Title, requestData.Description)

		c.JSON(http.StatusOK, gin.H{
			"redirect_url": "/ticket/" + encoded,
		})
	} else {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"errors": map[string]string{
					"title":       errors.Title,
					"description": errors.Description,
				},
				"values": map[string]string{
					"title":       requestData.Title,
					"description": requestData.Description,
				},
			},
		)
	}
}
