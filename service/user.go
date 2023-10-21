package service

import (
	"blog/config"
	daoutils "blog/dao/daoUtils"
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"
	"blog/response"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 注册逻辑处理
func Register(params *models.ParamsRegister) (err error) {
	if err := mysql.CheckUserExistByUsername(params.Username); err != nil {
		return err
	}

	newUser := &models.User{
		Uid:      snowflake.GenerateID(),
		Password: params.Password,
		Username: params.Username,
	}

	return mysql.CreateUser(newUser)
}

// 普通登录逻辑处理
func Login(params *models.ParamsLogin) (*LoginReply, error) {

	userInfo := new(models.User)
	if err := mysql.GetUserInfoByUsername(params.Username, userInfo); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(params.Password)); err != nil {
		return nil, response.CodeInvalidPassword
	}

	role, err := daoutils.GetRoleById(userInfo.RoleId)
	if err != nil && errors.Is(err, redis.ErrorRoleIsNotExist) {
		return nil, err
	}

	isAdmin := role != nil && role.IsAdmin == 1
	data, err := createReplyToken(userInfo.Uid, userInfo.RoleId, userInfo.Username, isAdmin)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// 刷新登录状态
func RefreshToken(claim *jwt.UserClaim) (*LoginReply, error) {
	return createReplyToken(claim.Uid, claim.RoleId, claim.Username, claim.IsAdmin)
}

func createReplyToken(uid int64, roleId, username string, isAdmin bool) (*LoginReply, error) {
	data := new(LoginReply)
	accessToken, err := jwt.GenerateToken(uid, config.Config.Jwt.AccessExpireAt, roleId, username, isAdmin)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(uid, config.Config.Jwt.RefreshExpireAt, roleId, username, isAdmin)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	data.AccessToken = accessToken
	data.RefreshToken = refreshToken
	data.AccessTokenExpireAt = now.Add(time.Duration(config.Config.Jwt.AccessExpireAt) * time.Hour).Format(time.DateTime)
	data.RefreshTokenExpireAt = now.Add(time.Duration(config.Config.Jwt.RefreshExpireAt) * time.Hour).Format(time.DateTime)

	return data, nil
}
