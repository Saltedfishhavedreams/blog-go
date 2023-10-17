package service

import (
	"blog/dao/mysql"
	"blog/models"
)

func Register(params *models.ParamRegister) (err error) {
	if err := mysql.CheckUserExistByUsername(params.Username); err != nil {
		return err
	}

	newUser := &models.User{
		Password: params.Password,
		Username: params.Username,
	}

	return mysql.CreateUser(newUser)
}
