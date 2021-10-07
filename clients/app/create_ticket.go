package app_client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/ticket-creator/usecases"
)

func (a *AppClient) CreateTicket(c *gin.Context) {
	c.Request.ParseForm()

	title := c.PostForm("title")
	description := c.PostForm("description")

	form := usecases.ValidateForm(
		title,
		description,
	)

	if form.Ok {
		encoded := usecases.EncodeTicketData(title, description)

		c.Redirect(
			http.StatusMovedPermanently,
			"/ticket/"+encoded,
		)
	} else {
		c.HTML(
			http.StatusUnprocessableEntity,
			"index.tmpl.html",
			map[string]interface{}{
				"Errors": map[string]string{
					"title":       form.Title,
					"description": form.Description,
				},
				"Values": map[string]string{
					"title":       title,
					"description": description,
				},
			},
		)
	}
}
