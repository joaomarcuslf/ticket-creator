package app_client

import "github.com/gin-gonic/gin"

type AppClient struct {
	Port string
}

func NewAppClient(port string) *AppClient {
	return &AppClient{
		Port: port,
	}
}

func (a *AppClient) Initialize() {
	router := gin.New()

	router.Use(gin.Logger())

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		a.Index(c)
	})

	router.GET("/ticket/:encodedUrl", func(c *gin.Context) {
		a.GetTicket(c)
	})

	router.POST("/create-ticket", func(c *gin.Context) {
		a.CreateTicket(c)
	})

	router.Run(":" + a.Port)
}
