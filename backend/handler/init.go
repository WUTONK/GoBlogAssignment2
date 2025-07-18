package handler

import (
	"WutonkGinBlog/models"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InitGin(g gin.IRouter) {
	g.POST("/user/login", Login)
	g.GET("/user/info", Info)
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
	token := md5.New()
	token.Write([]byte(username)) //写入并生成哈希
	// token.Sum(nil) 获取最终的字节数组（哈希值）
	return base64.URLEncoding.EncodeToString(token.Sum(nil))
}

func Login(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	valid := validUserAndPassword(req.Username, req.Password)

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "错误的用户名或密码"})
	}

	token := makeToken(req.Password)

	c.JSON(http.StatusOK, models.LoginResp{
		Token: makeToken(req.Password),
	})

	// 写入token
	if err := os.WriteFile("../tokenList/token.txt", []byte(token), 0666); err != nil {
		log.Fatal(err)
	}

}

func Info(c *gin.Context) {
	authorization := c.GetHeader("Authorization") // 获取头内容

	token, err := os.ReadFile("../tokenList/token.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token 文件读取失败"})
		return
	}
	if string(token) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "在token缓存中不存在任何token 无法验证"})
		fmt.Println("err:在token缓存中不存在任何token 无法验证")
		return
	}

	fmt.Printf("authorization:%s\n", authorization)
	fmt.Printf("token:%s\n", string(token))
	if authorization == string(token) {
		c.JSON(http.StatusOK, models.InfoResp{
			NikeName: "nnnnnikename!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "传入token与本地token不一致"})
		fmt.Println("err:传入token与本地token不一致")
	}

}
