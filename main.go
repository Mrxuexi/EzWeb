/*
@Time : 2021/8/16 下午4:01
@Author : mrxuexi
@File : main
@Software: GoLand
*/
package main

import (
	"Ez"
	"fmt"
	"net/http"
)
func main() {
	r := Ez.New()
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w,"Hello")
	})
	r.Run(":9090")
}
