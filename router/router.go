package router

import (
	"g/app/api"
	"g/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gsession"
	"time"
)

func init() {
	println("route.go init")
	s := g.Server()

	_ = s.SetConfigWithMap(g.Map{
		"SessionMaxAge":       time.Hour * 24,
		"SessionStorage":      gsession.NewStorageRedis(g.Redis()), // Redis存储Session
		"SessionCookieOutput": false,                               // 关闭自动设置Cookie功能
	})

	// 分组路由注册方式
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Session,
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.ALL("/captcha.html", api.Captcha.Get)
		group.ALL("/index/Login", api.Login)
		//group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.Auth)
			group.ALL("/user/profile", api.User.Profile)
		})
	})
}
