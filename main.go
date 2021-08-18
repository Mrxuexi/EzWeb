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

	r.POST("/hello/:id/*filepath", func(c *Ez.Context) {
		c.JSON(http.StatusOK,Ez.H{
			"name" : c.PostForm("name"),
			"age" : c.PostForm("age"),
			"id" : c.Param("id"),
			"filepath" : c.Param("filepath"),
		})
	})


	r.Run(":9090")
}
