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
	//给所有的路由组都添加了中间件logger
	r.Use(Ez.Logger())
	api := r.Group("/api")

	api.POST("/hello", func(c *Ez.Context) {
		c.JSON(http.StatusOK,Ez.H{
				"message" : "hello",
		})
	})
	//next的应用
	api.Use(func(c *Ez.Context) {
		c.JSON(200,Ez.H{
			"test" : "middleware2-1",
		})
		c.Next()
		c.JSON(200, Ez.H{
			"test" : "middleware2-2",
		})
	})

	r.Run(":9090")
}
