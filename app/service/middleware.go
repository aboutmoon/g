package service

import (
	"fmt"
	"g/app/model"
	"g/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"strings"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// Session中间件
func (s *middlewareService) Session(r *ghttp.Request) {

	token := fmt.Sprintf("%v", r.Get("token"))

	var sessionId string
	if token == "" {
		sessionId = r.Cookie.Get(r.Server.GetSessionIdName())
	} else {
		sessionId = token
	}

	if sessionId != "" {
		println("set id")
		_ = r.Session.SetId(sessionId)
	}

	// 执行下一步请求逻辑
	r.Middleware.Next()
	println("sessionid:" + r.Session.Id())
	r.Cookie.SetHttpCookie(&http.Cookie{
		Name:     r.Server.GetSessionIdName(),
		Value:    r.Session.Id(),
		Secure:   true,
		SameSite: http.SameSiteNoneMode, // 自定义samesite，配合secure一起使用
		Domain:   r.Server.GetCookieDomain(),
		Path:     r.Server.GetCookiePath(),
	})
}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 鉴权中间件，只有登录成功之后才能通过
func (s *middlewareService) Auth(r *ghttp.Request) {
	multiNoCheck := g.ArrayStr{
		"index/index/index",
		"index/index/navJson",
		"index/index/welcome",
		"index/index/alertSkin",
		"index/user/baseInfo",
		"index/login/login",
		"index/login/logout",
		"apps/game/test",
		"index/login/token",
	}

	// todo: 从配置读取
	systemKey := "app_gamesystem"

	path := strings.Replace(strings.ToLower(fmt.Sprintf("%v", r.URL.Path)), ".html", "", -1)
	path = strings.Replace(path, "/api/", "", -1)
	nodePermission := r.Session.GetStrings(systemKey+"nodePermissions", g.ArrayStr{})

	exist := false
	for _, v := range multiNoCheck {
		println(v, path)
		if v == path {
			exist = true
		}
	}

	if !exist {
		if len(nodePermission) <= 0 && path != "index/login/login" && path != "login/login" {
			response.JsonExit(r, 10007, "未登录！")
		} else {
			println("haha")
			r.Response.RedirectTo("/index/login/login")
		}
	}

	r.Middleware.Next()

	//r.Response.WriteStatus(http.StatusForbidden)
}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
