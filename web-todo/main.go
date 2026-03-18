package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.Run(":8080")
}
