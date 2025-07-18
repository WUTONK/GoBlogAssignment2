package handler

import (
	"WutonkGinBlog/models"
	"crypto/rand"
	"encoding/base64"
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

	valid, err := validUserAndPassword(req.Id, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证时后端发生错误，错误为: %s", err.Error())})
	}
	if valid {
		c.JSON(http.StatusOK, models.LoginResp{
			Token: makeToken(req.Id),
		})
	} else {

	}

}
