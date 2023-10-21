package models

type Role struct {
	Id           int    `json:"id" db:"id"`
	RoleId       string `json:"role_id" db:"role_id"`
	RoleName     string `json:"role_name" db:"role_name"`
	RoleNickname string `json:"role_nickname" db:"role_nickname"`
	RoleDesc     string `json:"role_desc" db:"role_description"`
	IsAdmin      uint8  `json:"is_admin" db:"is_admin"`
	CreateTime   string `json:"create_time" db:"create_time"`
	UpdateTime   string `json:"update_time" db:"update_time"`
}
