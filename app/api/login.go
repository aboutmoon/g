package api

import (
	"fmt"
	"g/app/model"
	"g/library/response"
	"github.com/dchest/captcha"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/techoner/gophp/serialize"
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

	//url := "http://test-permission.zxmn2018.com"
	//request := "/api/checklogin"

	system_key := "app_gamesystem"
	secret := "c2711e4866767ae07061578655ccd51b"
	param := g.MapStrStr{
		"user_name":  loginReq.Username,
		"password":   loginReq.Password,
		"ip":         r.GetClientIp(),
		"system_key": system_key,
	}

	p, err := serialize.Marshal(param)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	pi := fmt.Sprintf("%v", p)
	param["sign"], _ = gmd5.EncryptString(pi + secret)

	fmt.Printf("%+v %s", param, p)

	if !captcha.VerifyString(id, loginReq.Captcha) {
		response.JsonExit(r, 1, "验证码错误")
	}

	response.JsonExit(r, 1, "success", 1)
}
