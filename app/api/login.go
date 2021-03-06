package api

import (
	"bytes"
	"encoding/json"
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
	response.JsonExit(r, 1, "success", r.Session.Id())
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

	url := "http://test-permission.zxmn2018.com"
	loginRequest := "/api/checklogin"

	systemKey := "app_gamesystem"
	secret := "c2711e4866767ae07061578655ccd51b"
	paramMap := g.MapStrStr{
		"user_name":  loginReq.Username,
		"password":   loginReq.Password,
		"ip":         r.GetClientIp(),
		"system_key": systemKey,
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
		//response.JsonExit(r, 1, "验证码错误")
	}

	c, err := g.Client().Post(
		url+loginRequest,
		paramMap,
	)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		defer c.Close()
	}

	var res *model.PermissionRequestRes
	err = json.Unmarshal(c.ReadAll(), &res)
	if err != nil {
		fmt.Println(err)
		response.JsonExit(r, 1, err.Error())
	}

	// todo: 读配置
	_ = r.Session.Set("uid", res.Data.UserInfo.Id)
	_ = r.Session.Set("username", res.Data.UserInfo.UserName)
	_ = r.Session.Set("real_name", res.Data.UserInfo.UserName)
	_ = r.Session.Set(systemKey+"menu", res.Data.Menu)
	_ = r.Session.Set(systemKey+"nodePermissions", res.Data.NodePermissions)
	_ = r.Session.Set(systemKey+"platforms", res.Data.Platforms)
	_ = r.Session.Set(systemKey+"groups", res.Data.Groups)
	res.Data.SessionId = r.Session.Id()

	response.JsonExit(r, 1, "success", res.Data)
}
