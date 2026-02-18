//ausführen mit
// go run main.go

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	var website_name string = "Webseite von Cem"

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	group := r.Group("/resources")
	{
		group.GET("register", func(c *gin.Context) {
			c.HTML(200, "register.html", gin.H{
				"name": website_name,
			})
		})

		group.GET("login", func(c *gin.Context) {
			c.HTML(200, "login.html", gin.H{
				"name": website_name,
			})
		})
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"name": website_name,
		})
	})

	r.Run(":8080")
}
