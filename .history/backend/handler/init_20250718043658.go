package handler

import (
	"WutonkGinBlog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initGin(g gin.IRouter) {
	g.POST("/user/login", login)
}

func validUserAndPassword(user, password string) (bool, error) {
	return false
}

func login(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	valid, err := validUserAndPassword(req.Id, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证时后端发生错误，错误为: %s", err.Error())})
	}

}
