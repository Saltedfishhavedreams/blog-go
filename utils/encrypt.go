package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPwd(pwd string) string {
	newPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(newPwd)
}
