package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", helloRoute)
	}

	router.Run()
}

func helloRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World !",
		"status":  "done",
	})
}
