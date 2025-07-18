package main

import (
	"WutonkGinBlog/handler"

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

}
