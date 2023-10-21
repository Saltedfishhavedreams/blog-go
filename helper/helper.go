package helper

import (
	"math/rand"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(pwd string) string {
	newPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(newPwd)
}

func RandStr(length int) string {
	charSlice := make([]byte, length)
	char := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < length; i++ {
		charSlice[i] = char[rand.Intn(len(char))]
	}

	return string(charSlice)
}

func UUID() string {
	return uuid.NewV4().String()
}
