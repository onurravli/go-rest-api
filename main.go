package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/onurravli/go-rest-api/config"
)

func main() {
	cfg := config.GetConfig()
	var host = cfg.Api.Host
	var port = cfg.Api.Port

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	fmt.Printf("Listening on %s:%d", host, port)

	r.Run(fmt.Sprintf("%s:%d", host, port))

}
