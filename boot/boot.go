package boot

import (
	_ "g/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	println("boot.go init")
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
