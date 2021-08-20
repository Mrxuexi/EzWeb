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

	api := r.Group("/api")
	{
		api.POST("/hello", func(c *Ez.Context) {
			c.JSON(http.StatusOK,Ez.H{
				"message" : "hello",
			})
		})
		api.GET("/login", func(c *Ez.Context) {
			c.JSON(200,Ez.H{
				"name" : c.PostForm("name"),
			})
		})
	}


	r.Run(":9090")
}
