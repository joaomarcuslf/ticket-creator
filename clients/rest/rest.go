package rest_client

import "github.com/gin-gonic/gin"

type RestClient struct {
	Port string
}

func NewRestClient(port string) *RestClient {
	return &RestClient{
		Port: port,
	}
}

func (a *RestClient) Initialize() {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		a.Index(c)
	})

	router.GET("/ticket/:encodedUrl", func(c *gin.Context) {
		a.GetTicket(c)
	})

	router.POST("/ticket", func(c *gin.Context) {
		a.CreateTicket(c)
	})

	router.Run(":" + a.Port)
}
