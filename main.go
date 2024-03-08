package main

import (
	"Egin"
	"net/http"
)

func main() {
	r := Egin.New()
	r.GET("/", func(c *Egin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *Egin.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *Egin.Context) {
		c.JSON(http.StatusOK, Egin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
