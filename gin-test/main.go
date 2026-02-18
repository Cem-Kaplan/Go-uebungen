package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gruppe := r.Group("/main")
	{
		gruppe.GET("/login", func(c *gin.Context) {
			c.String(200, "Test: login")
		})

		gruppe.GET("/register", func(c *gin.Context) {
			c.String(200, "Test: register")
		})
	}
	r.Run(":8080")
}
