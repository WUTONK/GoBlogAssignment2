package handler

import (
	"WutonkGinBlog/models"

	"github.com/gin-gonic/gin"
)

func initGin(g gin.IRouter) {
	g.POST("/user/login", login)
}

func validUserAndPassword(user, password string) {

}

func login(c *gin.Context) {
	var req models.LoginReq

}
