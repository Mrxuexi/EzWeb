/*
@Time : 2021/8/16 下午4:03
@Author : Mrxuexi
@File : Ez
@Software: GoLand
*/

package Ez

import (
	"net/http"
)

// HandlerFunc 是Ez框架中定义的对请求的响应处理方法，传入*Context针对http请求处理
type HandlerFunc func(*Context)

// Engine 实现了"net/http"标准库中的 Handler 接口中的ServeHTTP方法
type Engine struct {
	//用于存储路由处理方法
	//key是方法类型加路径，value是用户的处理方法
	router *router
}

// ServeHTTP 方法的实现，用于实现处理HTTP请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//根据req和w实例一个context
	c := newContext(w, req)
	//传入开始执行处理
	engine.router.handle(c)
}

// New 路由存储结构的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 方法封装在router中，在 router map[string]HandlerFunc 中存入对应处理方法
func (engine *Engine) addRoute(method string, path string, handler HandlerFunc) {
	engine.router.addRoute(method, path, handler)
}

// GET 实现的是注册GET请求的路径和对应方法，调用了addRoute，存入了route 结构体的handler中
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