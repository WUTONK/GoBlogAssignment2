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

func InitGin(g gin.IRouter) {
	g.POST("/user/login", login)
	g.GET("/user/info", info)
}

func validUserAndPassword(incomingUsername, incomingPassword string) bool {

	fmt.Printf("收到的用户名%s\n", incomingUsername)
	fmt.Printf("收到的密码%s\n", incomingPassword)
	// 密码123456的密文
	passwordHash := "$2a$10$0y3/DmaCYAIRrcPM52TJ6.S/ax2nkl9Avbfp2XS.rATqvWFNIy58G%!"
	username := "kevin"

	usernameValid := false
	passwordValidErr := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(incomingPassword))

	if incomingUsername == username {
		usernameValid = true
		fmt.Println("用户名验证通过")
	}

	if usernameValid && passwordValidErr == nil {
		return true
	} else {
		return false
	}
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

	valid := validUserAndPassword(req.Username, req.Password)

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

func info(c *gin.Context) {
	c.String(200, "hi kk")
}
