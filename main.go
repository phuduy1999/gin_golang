package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routingGet(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {

		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	u1 := router.Group("/user")
	{
		u1.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		u1.GET("/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})

		u1.GET("/groups", func(c *gin.Context) {
			c.String(http.StatusOK, "The available groups are [...]")
		})
	}
}

func routingPost(router *gin.Engine) {

	router.POST("/user/:name/*action", func(c *gin.Context) {
		if c.FullPath() == "/user/:name/*action" {
			c.String(http.StatusOK, "ddsa")
		} // true
	})

	// router.POST("/form_post", handlers.formPost())

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		// c.JSON(200, gin.H{
		// 	"id":      id,
		// 	"page":    page,
		// 	"name":    name,
		// 	"message": message,
		// })
		c.String(200, fmt.Sprintf("id: %s; page: %s; name: %s; message: %s", id, page, name, message))
	})

	router.POST("/posts", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.String(200, fmt.Sprintf("ids: %v; names: %v", ids, names))
	})
}

func main() {
	router := gin.Default()

	routingGet(router)
	routingPost(router)

	router.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
