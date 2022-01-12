package api

import (
	"bytes"
	"fmt"
	"g/app/model"
	"g/library/response"
	"github.com/dchest/captcha"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/techoner/gophp/serialize"
	"sort"
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
	paramMap := g.MapStrStr{
		"user_name":  loginReq.Username,
		"password":   loginReq.Password,
		"ip":         r.GetClientIp(),
		"system_key": system_key,
	}

	paramArr := g.ArrayStr{}
	for _, v := range paramMap {
		paramArr = append(paramArr, v)
	}

	sort.Strings(paramArr)

	p, err := serialize.Marshal(paramArr)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	p = bytes.Join([][]byte{p, []byte(secret)}, []byte(""))
	paramMap["sign"], _ = gmd5.EncryptBytes(p)

	fmt.Printf("%+v %s", paramMap, p)

	if !captcha.VerifyString(id, loginReq.Captcha) {
		response.JsonExit(r, 1, "验证码错误")
	}

	response.JsonExit(r, 1, "success", 1)
}
