package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建 Gin Engine 实例
	r := gin.Default()

	// 设置请求 URI /ping 的路由及响应处理函数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 启动 Web 服务，监听端口,等待 HTTP 请求到并生成响应
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
