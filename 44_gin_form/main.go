package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
