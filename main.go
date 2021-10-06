package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	handlers "github.com/joaomarcuslf/ticket-creator/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()

	router.Use(gin.Logger())

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		handlers.Index(c)
	})

	router.GET("/ticket/:encodedUrl", func(c *gin.Context) {
		handlers.GetTicket(c)
	})

	router.POST("/create-ticket", func(c *gin.Context) {
		handlers.CreateTicket(c)
	})

	router.Run(":" + port)
}
