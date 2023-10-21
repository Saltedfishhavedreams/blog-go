package service

import "blog/models"

type LoginReply struct {
	AccessToken          string `json:"access_token"`
	RefreshToken         string `json:"refresh_token"`
	AccessTokenExpireAt  string `json:"access_token_expire_at"`
	RefreshTokenExpireAt string `json:"refresh_token_expire_at"`
}

type RoleReply struct {
	models.Role
}
