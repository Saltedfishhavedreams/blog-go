package jwt

import (
	"blog/config"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Uid      int64  `json:"uid"`
	RoleId   string `json:"role_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"` // 是否是超管
	jwt.StandardClaims
}

// 生成token
func GenerateToken(uid int64, expireAt int, roleId, username string, isAdmin bool) (string, error) {
	currUser := UserClaim{
		Uid:      uid,
		RoleId:   roleId,
		Username: username,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expireAt) * time.Hour).Unix(),
			Issuer:    config.Config.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, currUser)
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.Secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func AnalyzeToken(tokenString string) (*UserClaim, error) {
	userClaim := new(UserClaim)
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return userClaim, nil
	}
	return nil, errors.New("invalid token")
}
