package models

type User struct {
	Id                int    `json:"id" db:"id"`
	Uid               int64  `json:"uid" db:"uid"`
	Username          string `json:"username" db:"username"`
	Password          string `json:"password" db:"password"`
	Email             string `json:"email" db:"email"`
	Avatar            string `json:"avatar" db:"avatar"`
	Gender            int16  `json:"gender" db:"gender"`
	PersonalSignature string `json:"personal_signature" db:"personal_signature"`
	CreateTime        string `json:"create_time" db:"create_time"`
	UpdateTime        string `json:"update_time" db:"update_time"`
	Nickame           string `json:"nickname" db:"nickname"`
	RoleId            string `json:"role_id" db:"role_id"`
}
