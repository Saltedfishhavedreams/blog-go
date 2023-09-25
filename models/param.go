package models

// 账号注册
type ParamRegister struct {
	Username   string `json:"username" binding:"required,max=24,min=6"`
	Password   string `json:"password" binding:"required,max=24,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
