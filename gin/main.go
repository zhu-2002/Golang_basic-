package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// gin就是Engine实例，负责路由注册、接入middleware的核心职责
	r := gin.Default()
	// 路由注册，主要参数是指定路由规则（/hello）和处理函数
	// gin.Context核心职责是处理请求和返回响应
	// 静态路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个静态路由")
	})
	// 参数路由
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "这是一个参数：%s", name)
	})
	// 通配符路由
	r.GET("/view/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(http.StatusOK, "匹配参数：%s", path)
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
