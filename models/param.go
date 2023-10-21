package models

// 账号注册
type ParamsRegister struct {
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	ParamsLogin
}

// 普通登录
type ParamsLogin struct {
	Username string `json:"username" binding:"required,max=24,min=6"`
	Password string `json:"password" binding:"required,max=24,min=6"`
}

// 创建角色
type ParamsCreateRole struct {
	RoleName     string `json:"role_name" binding:"required,max=24,min=6"`
	RoleNickname string `json:"role_nickname" binding:"required,max=24,min=6"`
	RoleDesc     string `json:"role_desc" binding:"max=200,min=0"`
}
