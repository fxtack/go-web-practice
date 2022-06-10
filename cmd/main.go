package main

import (
	"fmt"
	"go-web-practice/internal/config"
	"go-web-practice/internal/controller"
	"go-web-practice/internal/log"
	"go-web-practice/internal/middleware"
	"go-web-practice/internal/templates"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {

	// 加载 json 配置
	config.LoadConfig("configs/config.json")

	// 加载日志模块
	log.LoadLog()

	// 载入模板
	templates.LoadTemplate(config.Config.Template)

	// 设置服务
	server := http.Server{
		Addr: config.Config.Address + ":" + strconv.Itoa(config.Config.Port),
		Handler: &middleware.TimeoutMiddleWare{
			Next: &middleware.LogMiddleWare{
				Next: &middleware.CrossMiddleWare{},
			},
		},
	}

	// 注册路由
	controller.RegisterRoutes()

	// 显示启动信息
	fmt.Printf("\x1b[1;32m* Web services run on:\x1b[0m\t"+
		"http://%s:%d\n", "localhost", config.Config.Port)


	server.ListenAndServe()
}
