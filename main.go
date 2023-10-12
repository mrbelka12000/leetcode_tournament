package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from zeet, The project was deployed",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
	// listen and serve on 0.0.0.0:8080
}
