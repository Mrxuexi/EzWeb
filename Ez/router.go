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
	"strings"
)

type router struct {
	//用于存储相关方法
	handlers map[string]HandlerFunc
	//用于存储每种请求方式的树的根节点，用于辅助处理动态路由
	roots map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots: make(map[string]*node),
	}
}

// parsePath 用于处理传入的url，先将其分开存储到parts中，当然直到出现*前缀的部分，就可以结束
func parsePath(path string) []string {
	vs := strings.Split(path, "/")
	parts := make([]string, 0)
	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}

// router 中 addRoute 方法，在 handlers map[string]HandlerFunc 中存入路由对应处理方法
//存入形式为例如：{ "GET-/index" : 定义的处理方法 }
func (r *router) addRoute(method string, path string, handler HandlerFunc)  {
	parts := parsePath(path)
	log.Printf("Route %4s  -  %s",method,path)
	key := method + "-" + path
	_, ok := r.roots[method]
	//roots中不存在对应的方法入口则注册相应方法入口
	if !ok {
		r.roots[method] = &node{}
	}
	//调用路由表插入方法，在该数据结构中插入该路由
	r.roots[method].insert(path, parts, 0)
	//把method-path作为key，以及handler方法作为value注入数据结构
	r.handlers[key] = handler
}

//获得路由，先将传入的路径字符串处理成字符串数组,然后根据method进入到对应的路由树的入口
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePath(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)//传入全部路径的字符串数组，寻找到最后对应节点

	if n != nil {
		parts := parsePath(n.path) //n.path包含了完整的路由

		for i, part := range parts {//遍历这一条路径
			//拿到:的参数，存入params，方法中的part作为key，外面传入的path中的数据作为value存入
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
			//拿到*，此时路由表中的存入的part作为key,外面传入的path中的数据作为value传入params，之后也再没有了
			if part[0] == '*' && len(part) > 1{
				params[part[1:]] = strings.Join(searchParts[i:],"/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

//根据context中存储的 c.Method 和 c.Path 拿到对应的处理方法，进行执行，如果拿到的路由没有注册，则返回404
func (r *router) handle(c *Context)  {
	//获取匹配到的节点，同时也拿到两类动态路由中参数
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		//拿目的节点中的path做key来找handlers
		key := c.Method + "-" + n.path
		//根据路径拿到处理器
		c.handlers = append(c.handlers, r.handlers[key])
	}else {
		//不存在节点的情况下，给生成的c加入一个404方法
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: ", c.Path)
		})
	}
	c.Next()
}