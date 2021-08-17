/*
@Time : 2021/8/17 下午1:46
@Author : Mrxuexi
@File : context.go
@Software: GoLand
*/
package Ez

import (
	"encoding/json"
	"net/http"
)

// H 为map[string]interface{}结构体起个别名，用户在代码中构建JSON时显得更简洁
type H map[string]interface{}

// Context 结构体，内部封装了 http.ResponseWriter, *http.Request
type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	//请求的信息，包括路由和方法
	Path string
	Method string
	//响应的状态码
	StatusCode int
}

//Context构造方法
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:     w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
	}
}

// 访问参数的处理方法PostForm和Query

// PostForm 根据key拿到请求中的表单内容
func (c Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 根据key获取请求中的参数
func (c Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//一些服用的前值方法Status处理响应状态码 SetHeader处理响应消息头

// Status 将状态码写入context，同时将通过封装起来的http.ResponseWriter方法，将状态码写入响应头
func (c Context) Status(code int)  {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 构造响应的消息头
func (c Context) SetHeader(key string,value string)  {
	c.Writer.Header().Set(key,value)
}

// String 调用我们的 SetHeader和Status 方法，构造string类型响应的状态码和消息头，然后将字符串转换成byte写入到响应头
func (c Context) String(code int,values ...interface{})  {
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	var str = ""
	for _, value := range values {
		str += value.(string)
	}
	c.Writer.Write([]byte(str))
}

// JSON 调用我们的 SetHeader和Status 方法，构造JSON类型响应的状态码和消息头，根据我们传入的对象来构造json数据写入
func (c Context) JSON(code int,obj interface{})  {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(),http.StatusInternalServerError)
	}
}

// Data 同上 ，但是直接写入字节数组，不再构建消息头
func (c Context) Data(code int,data []byte)  {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 模版渲染 同上，消息体传入的是html文件
func (c Context) HTML(code int, html string)  {
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}