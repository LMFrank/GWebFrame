package main

import (
	"net/http"

	"gwebframe"
)

func main() {
	r := gwebframe.New()
	r.GET("/", func(c *gwebframe.Context) {
		c.HTML(http.StatusOK, "<h1>Hello world</h1>")
	})

	r.GET("/hello", func(c *gwebframe.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gwebframe.Context) {
		c.JSON(http.StatusOK, gwebframe.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8000")
}