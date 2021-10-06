package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	fmt.Println("The URL: ", c.Request.Host+c.Request.URL.Path)

	requestBody := strings.NewReader(`
	{
		"long_url": "` + c.Request.Host + c.Request.URL.Path + `"
	}
	`)

	// post some data
	res, err := http.Post(
		"https://vast-thicket-55540.herokuapp.com/create-short-url",
		"application/json; charset=UTF-8",
		requestBody,
	)

	// check for response error
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", data)

	var raw map[string]interface{}

	if err := json.Unmarshal(data, &raw); err != nil {
		panic(err)
	}

	c.HTML(
		http.StatusOK,
		"ticket.tmpl.html",
		map[string]interface{}{
			"Values": map[string]interface{}{
				"title":            s[0],
				"description":      description,
				"description_safe": s[1],
				"date":             "2021-10-06",
				"short_url":        raw["short_url"],
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
