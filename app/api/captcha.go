package api

import (
	"g/library/response"
	"github.com/dchest/captcha"
	"github.com/gogf/gf/net/ghttp"
)

var Captcha = new(captchaApi)

type captchaApi struct{}

// 用户注册
func (s *captchaApi) Get(r *ghttp.Request) {
	r.Response.ResponseWriter.Header().Set("Content-Type", "image/png; charset=utf-8")
	id := captcha.NewLen(4)

	err := r.Session.Set("captcha", id)
	if err != nil {
		response.JsonExit(r, 1, "设置session失败")
	}
	_ = captcha.WriteImage(r.Response.Writer, id, 200, 100)
	//img := captcha.NewImage("123", captcha.RandomDigits(5), 100, 100)
	//img.WriteTo(r.Response.Writer)
}
