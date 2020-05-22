package main

import (
	"net/http"

	"gwebframe"
)

func main() {
	r := gwebframe.New()
	r.GET("/index", func(c *gwebframe.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gwebframe.Context) {
			c.HTML(http.StatusOK, "<h1>Hello gwebframe</h1>")
		})

		v1.GET("/hello", func(c *gwebframe.Context) {
			// /hello?name=admin
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gwebframe.Context) {
			// /hello/admin
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gwebframe.Context) {
			c.JSON(http.StatusOK, gwebframe.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":8000")
}
