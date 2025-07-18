package handler

import (
	"WutonkGinBlog/models"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func initGin(g gin.IRouter) {
	g.POST("/user/login", login)
}

func validUserAndPassword(user, password string) bool {
	// 密码123456的密文：$2a$10$0y3/DmaCYAIRrcPM52TJ6.S/ax2nkl9Avbfp2XS.rATqvWFNIy58G%!

	passwordHash := "$2a$10$0y3/DmaCYAIRrcPM52TJ6.S/ax2nkl9Avbfp2XS.rATqvWFNIy58G%!"

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		fmt.Printf(err.Error())
		return false
	} else {
		return true
	}

	return false
}

func makeToken(username string) string {
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

	valid := validUserAndPassword(req.Password, req.Password)

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "错误的用户名或密码"})
	}

	token := makeToken(req.Password)

	if token == "error" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "生成token时发生错误"})
	}

	c.JSON(http.StatusOK, models.LoginResp{
		Token: makeToken(req.Password),
	})
}
