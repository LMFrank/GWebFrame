package main

import (
	"fmt"
	"gwebframe"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gwebframe.Default()
	r.GET("/", func(c *gwebframe.Context) {
		c.String(http.StatusOK, "Hello world\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gwebframe.Context) {
		names := []string{"admin"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8000")
}

//func main() {
//	r := gwebframe.New()
//	r.Use(gwebframe.Logger())
//	r.SetFuncMap(template.FuncMap{
//		"formatAsDate": formatAsDate,
//	})
//	r.LoadHTMLGlob("templates/*")
//	r.Static("/assets", "./static")
//
//	stu1 := &student{Name: "Tom", Age: 20}
//	stu2 := &student{Name: "Jack", Age: 22}
//	r.GET("/", func(c *gwebframe.Context) {
//		c.HTML(http.StatusOK, "css.tmpl", nil)
//	})
//	r.GET("/students", func(c *gwebframe.Context) {
//		c.HTML(http.StatusOK, "arr.tmpl", gwebframe.H{
//			"title":  "gwebframe",
//			"stuArr": [2]*student{stu1, stu2},
//		})
//	})
//
//	r.GET("/date", func(c *gwebframe.Context) {
//		c.HTML(http.StatusOK, "custom_func.tmpl", gwebframe.H{
//			"title": "gwebframe",
//			"now":   time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
//		})
//	})
//
//	r.Run(":8000")
//}
