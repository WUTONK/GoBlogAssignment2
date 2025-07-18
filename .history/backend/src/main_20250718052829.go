package main

import (
	passwordhash "WutonkGinBlog/passwordHash"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Printf(passwordhash.PasswordHash("123456"))

	incomingPassword := "123456"
	passwordHash := "$2a$10$0y3/DmaCYAIRrcPM52TJ6.S/ax2nkl9Avbfp2XS.rATqvWFNIy58G%!"
	passwordValidErr := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(incomingPassword))
	fmt.Printf(passwordValidErr.Error())
	if passwordValidErr == nil {
		fmt.Printf("error is emtpy")
	}
}
