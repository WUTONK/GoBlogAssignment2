package main

import (
	passwordhash "WutonkGinBlog/passwordHash"
	"fmt"
)

func main() {
	fmt.Printf(passwordhash.PasswordHash("123456"))
}
