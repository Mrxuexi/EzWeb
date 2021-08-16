/*
@Time : 2021/8/16 下午4:03
@Author : Mrxuexi
@File : Ez
@Software: GoLand
*/

package Ez

import (
	"fmt"
	"net/http"
)

// HandlerFunc 是Ez框架中定义的对请求的响应处理方法，默认传入这两个参数,针对http请求处理
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现了"net/http"标准库中的 Handler 接口中的ServeHTTP方法
type Engine struct {
	//用于存储路由处理方法
	//key是方法类型加路径，value是用户的处理方法
	router map[string]HandlerFunc
}

// ServeHTTP 方法的实现，用于实现处理HTTP请求
// 先解析req对应传入的路径，查找router中，如果有相应的处理方法，则执行处理方法，如果没有则返回找不到的提示
// 来自Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	//根据请求req中的数据，从router中取出对应的方法
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "could not find the route: %s\n", req.URL)
	}
}

// New 是Ez.Engine的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// Engine 中 addRoute 方法，在 router map[string]HandlerFunc 中存入对应处理方法
//存入形式为例如：{ "GET-/index" : 定义的处理方法 }engine
func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	engine.router[key] = handler
}

// GET 实现的是Engine的处理GET请求的方法
func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.addRoute("GET", path, handler)
}

// POST 同上
func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.addRoute("POST", path, handler)
}


func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}