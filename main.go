package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 Gin 实例
	r := gin.Default()

	// 加载 HTML 模板文件
	r.LoadHTMLGlob("templates/*")

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		// 渲染模板并传递数据
		c.HTML(200, "index.html", gin.H{
			"title":   "选课社区",
			"message": "一个超级无敌大饼",
		})
	})

	// 启动服务器
	r.Run(":8080")
}
