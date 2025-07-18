package passwordhash

import "golang.org/x/crypto/bcrypt"

func main() {
	bytes, err := bcrypt.GenerateFrompwd([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
