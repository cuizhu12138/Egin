package Egin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 为map起了别名，构造时代码更简洁
type H map[string]interface{}

type Context struct {
	// 原有的组件
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求相关组件
	Path   string
	Method string
	// 应答相关组件
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 从POST头中解析对应参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 从URL中解析GET参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置响应的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置http头中的字段，key-value对
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 构造字符串，设置数据格式和状态码，最后把字符串输入
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 构造JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

// 设置Data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 快速构造HTML回应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

