package main

import (
	"backend/routers"
)

func main() {
	r := routers.InitRoutes()

	// 启动服务
	r.Run(":8080") // 默认监听 http://localhost:8080
}
