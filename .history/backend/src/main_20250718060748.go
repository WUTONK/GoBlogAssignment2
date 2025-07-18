package main

import (
	"WutonkGinBlog/handler"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "WUTONK")
	})

	handler.InitGin(router)

	log.Println("Server is running on port http://127.0.0.1:8080")

	// 服务器运行端口
	router.Run("127.0.0.1:8080")

}
