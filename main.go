/*
@Time : 2021/8/16 下午4:01
@Author : mrxuexi
@File : main
@Software: GoLand
*/
package main

import (
	"Ez"
	"net/http"
)
func main() {
	r := Ez.New()
	r.GET("/", func(c *Ez.Context) {
		c.HTML(http.StatusOK,"<h1>This is the index</h1>")
	})
	r.GET("/hello", func(c *Ez.Context) {
		c.String(http.StatusOK, "hello")
	})
	r.POST("/hello", func(c *Ez.Context) {
		c.JSON(http.StatusOK,Ez.H{
			"name" : c.PostForm("name"),
			"age" : c.PostForm("age"),
		})
	})

	r.Run(":9090")
}
