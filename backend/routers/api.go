package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// 提供静态文件服务
	router.Static("/static", "./static")

	// 提供模板文件
	router.LoadHTMLGlob("templates/*")

	// 路由组：API
	api := router.Group("/api")
	{
		api.GET("/courses", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"courses": []string{"数学", "物理", "化学"},
			})
		})
	}

	return router
}
