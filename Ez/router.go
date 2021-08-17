/*
@Time : 2021/8/17 下午3:31
@Author : Mrxuexi
@File : router
@Software: GoLand
*/
package Ez

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// router 中 addRoute 方法，在 handlers map[string]HandlerFunc 中存入路由对应处理方法
//存入形式为例如：{ "GET-/index" : 定义的处理方法 }
func (r *router) addRoute(method string, path string, handler HandlerFunc)  {
	log.Printf("Route %4s  -  %s",method,path)
	key := method + "-" + path
	r.handlers[key] = handler
}

//根据context中存储的 c.Method 和 c.Path 拿到对应的处理方法，进行执行，如果拿到的路由没有注册，则返回404
func (r *router) handle(c *Context)  {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	}else {
		c.String(http.StatusNotFound,"404 NOT FOUND")
	}
}