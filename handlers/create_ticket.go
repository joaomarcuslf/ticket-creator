package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/ticket-creator/usecases"
)

func CreateTicket(c *gin.Context) {
	c.Request.ParseForm()

	title := c.PostForm("title")
	description := c.PostForm("description")

	errors := usecases.ValidateForm(
		title,
		description,
	)

	if errors.Ok {
		encoded := usecases.EncodeTicketData(title, description)

		c.Redirect(
			http.StatusMovedPermanently,
			"/ticket/"+encoded,
		)
	} else {
		fmt.Print()
		c.HTML(
			http.StatusUnprocessableEntity,
			"index.tmpl.html",
			map[string]interface{}{
				"Errors": map[string]string{
					"title":       errors.Title,
					"description": errors.Description,
				},
				"Values": map[string]string{
					"title":       title,
					"description": description,
				},
			},
		)
	}
}
