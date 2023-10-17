package utils

import (
	"math/rand"
)

func RandStr(length int) string {
	charSlice := make([]byte, length)
	char := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < length; i++ {
		charSlice[i] = char[rand.Intn(len(char))]
	}

	return string(charSlice)
}
