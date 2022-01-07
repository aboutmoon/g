package api

import (
	"g/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var Login = new(loginApi)

type loginApi struct{}

// 获取token
func (a *loginApi) Token(r *ghttp.Request) {
	response.JsonExit(r, 1, "success", r.GetSessionId())
}
