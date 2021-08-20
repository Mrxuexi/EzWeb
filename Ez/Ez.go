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

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 用于存储分组的中间件
	engine      *Engine       // 这里实现的嵌套让其拥有了engine的全部属性,所有的分组都使用了Engine实例,可以通过engine间接的访问各种接口
}

// Engine 实现了"net/http"标准库中的 Handler 接口中的ServeHTTP方法
type Engine struct {
	*RouterGroup	//嵌套，让Engine拥有RouterGroup的全部属性，这样做体现在我们使用r.Group()创建路由组的时候
	groups []*RouterGroup
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

// New Engine
func New() *Engine {
	//新建一个engine实例
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 新建Group
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix:      group.prefix + prefix,		//前缀
		engine:      engine,					//任何路由组都共享一个处理实例
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// engine.router.addRoute 方法封装在router中
func (group *RouterGroup) addRoute(method string, last string, handler HandlerFunc) {
	path := group.prefix + last
	group.engine.router.addRoute(method, path, handler)
}

func (group *RouterGroup) GET(path string, handler HandlerFunc) {
	group.addRoute("GET", path, handler)
}

func (group *RouterGroup) POST(path string, handler HandlerFunc) {
	group.addRoute("POST", path, handler)
}


func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}