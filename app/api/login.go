package api

import (
	"fmt"
	"g/app/model"
	"g/library/response"
	"github.com/dchest/captcha"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var Login = new(loginApi)

type loginApi struct {
}

// 获取token
func (a *loginApi) Token(r *ghttp.Request) {
	response.JsonExit(r, 1, "success", r.GetSessionId())
}

//
func (a *loginApi) Login(r *ghttp.Request) {

	var (
		loginReq *model.UserLoginReq
	)

	if err := r.ParseForm(&loginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	id := fmt.Sprintf("%v", r.Session.Get("captcha"))
	if !captcha.VerifyString(id, loginReq.Captcha) {
		response.JsonExit(r, 1, "验证码错误")
	}

	response.JsonExit(r, 1, "success", 1)
}
