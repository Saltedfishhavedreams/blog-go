package mysql

import (
	"blog/models"
	"blog/utils"
	"database/sql"
	"fmt"
)

// @params conditionField 查询条件字段 (id 或 username)
// @params condition 查询条件值
func getUserInfo(conditionField, condition string) (*models.User, error) {
	switch condition {
	case "id":
	case "username":
		break
	default:
		conditionField = "id"
	}

	sqlStr := fmt.Sprintf("select id, avatar, gender, nickname, nickname, role_id, create_time, update_time from user where %s = ?", conditionField)
	userinfo := new(models.User)

	err := db.Get(userinfo, sqlStr, condition, condition)
	if err == sql.ErrNoRows {
		return nil, ErrorUserNotExist
	} else if err != nil {
		return nil, err
	}

	return userinfo, nil
}

// @params condition 查询条件值
func GetUserInfoById(condition string) (*models.User, error) {
	return getUserInfo("id", condition)
}

// @params condition 查询条件值
func GetUserInfoByUsername(condition string) (*models.User, error) {
	return getUserInfo("username", condition)
}

// @params username 用户名
func CheckUserExistByUsername(username string) error {
	sqlStr := "select count(id) from user where username = ?"
	var count int

	err := db.Get(&count, sqlStr, username)

	if err != nil {
		return err
	}

	if count > 0 {
		return ErrorUserExist
	}

	return nil
}

// @params otherFields 可选参数 [role_id, sort]
func CreateUser(user *models.User, otherFields ...interface{}) (err error) {
	sqlStr := "insert into user(username, password, role_id, sort, nickname) values(?, ?, ?, ?, ?)"
	randStr := utils.RandStr(15)

	var role_id int
	if len(otherFields) > 1 {
		if _role_id, ok := otherFields[0].(int); ok {
			role_id = _role_id
		}
	}

	var sort int
	if len(otherFields) > 2 {
		if _sort, ok := otherFields[1].(int); ok {
			sort = _sort
		}
	}

	nickame := randStr[:6] + "_" + randStr[6:]
	fmt.Printf("nickame: %v\n", nickame)

	_, err = db.Exec(sqlStr, user.Username, utils.EncryptPwd(user.Password), role_id, sort, nickame)

	return
}
