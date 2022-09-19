package main

import (
	"wechat/conf"
	"wechat/router"
)

func main() {

	// 初始化公众号配置
	if err := conf.InitConf(); err != nil {
		panic(err)
	}

	// 初始化gin路由
	r := router.InitRouter()
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
