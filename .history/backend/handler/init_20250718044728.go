package handler

import (
	"WutonkGinBlog/models"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initGin(g gin.IRouter) {
	g.POST("/user/login", login)
}

func validUserAndPassword(user, password string) bool {
	return false
}

func makeToken(Id string) string {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "error"
	}
	return base64.URLEncoding.EncodeToString(token)
}

func login(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	valid := validUserAndPassword(req.Id, req.Password)

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "错误的用户名或密码"})
	}

	Token := makeToken(req.Id)

	c.JSON(http.StatusOK, models.LoginResp{
		Token: makeToken(req.Id),
	})
}
