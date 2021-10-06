package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/ticket-creator/encoders"
)

func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		nil,
	)
}

func GetTicket(c *gin.Context) {
	encodedUrl := c.Param("encodedUrl")

	decodedUrl, _ := encoders.Decode(encodedUrl)

	s := strings.Split(decodedUrl, "-")

	description := strings.Split(s[1], "\r\n")

	c.HTML(
		http.StatusOK,
		"ticket.tmpl.html",
		map[string]interface{}{
			"Values": map[string]interface{}{
				"title":       s[0],
				"description": description,
				"date":        "2021-10-06",
			},
		},
	)
}

func CreateTicket(c *gin.Context) {
	c.Request.ParseForm()

	Errors := map[string]string{
		"title":       "",
		"description": "",
	}

	if c.PostForm("title") == "" {
		Errors["title"] = "Required field"
	} else if len(c.PostForm("title")) > 100 {
		Errors["title"] = "Field has max length of 100"
	}

	if c.PostForm("description") == "" {
		Errors["description"] = "Required field"
	}

	if Errors["title"] == "" && Errors["description"] == "" {
		encoded := encoders.Encode(
			fmt.Sprintf("%v-%v", c.PostForm("title"), c.PostForm("description")),
		)

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
				"Errors": Errors,
				"Values": map[string]string{
					"title":       c.PostForm("title"),
					"description": c.PostForm("description"),
				},
			},
		)
	}
}
