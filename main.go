package main

import (
	"log"
	"net/http"
	"time"

	"gwebframe"
)

func onlyForV2() gwebframe.HandlerFunc {
	return func(c *gwebframe.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gwebframe.New()
	r.Use(gwebframe.Logger()) // global midlleware
	r.GET("/", func(c *gwebframe.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gwebframe</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gwebframe.Context) {
			// expect /hello/admin
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":8000")
}
