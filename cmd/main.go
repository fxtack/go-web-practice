package main

import (
	"fmt"
	"go-web-practic/internal/config"
	"go-web-practic/internal/controller"
	"go-web-practic/internal/log"
	"go-web-practic/internal/middleware"
	"go-web-practic/internal/templates"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {

	// 加载 json 配置
	config.LoadConfig("../configs/config.json")

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
	fmt.Printf("\x1b[1;32m* The Web Server was running on:\x1b[0m\t"+
		"http://%s:%d\n", "localhost", config.Config.Port)
	fmt.Printf("\x1b[1;32m* The profile situation can be view on:\x1b[0m\t"+
		"http://localhost:%d/debug/pprof\n", config.Config.PprofPort)

	// 启动性能监控服务与 web 服务
	go http.ListenAndServe(":"+strconv.Itoa(config.Config.PprofPort), nil)
	server.ListenAndServe()
}
