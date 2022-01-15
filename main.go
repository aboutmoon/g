package main

import (
	_ "g/boot"
	_ "g/router"
	"github.com/gogf/gf/frame/g"
)

// @title       `gf-demo`示例服务API
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
func main() {
	println("main.go")
	g.Server().Run()
}
