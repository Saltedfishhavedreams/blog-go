package mysql

import (
	"blog/helper"
	"blog/models"
	"blog/response"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// @Description 根据 struct tag 中 DB 查询用户信息。
// @params conditionField 查询条件字段 (id 或 username)。
// @params condition 查询条件值。
// @params data 数据接收参数 (指针类型)。
func getUserInfo(conditionField, condition string, data any) error {
	rvData := reflect.ValueOf(data)
	rtData := rvData.Type()
	if rtData.Kind() == reflect.Ptr && !rvData.IsNil() {
		rtData = rtData.Elem()

		if rtData.Kind() == reflect.Struct {
			switch conditionField {
			case "id":
			case "username":
				break
			default:
				conditionField = "id"
			}

			// 获取所有需要查询的字段
			selectFields := make([]byte, 0)
			for fieldIndex := 0; fieldIndex < rtData.NumField(); fieldIndex++ {
				currField := rtData.Field(fieldIndex)
				currDBField := currField.Tag.Get("db")
				if currDBField != "" {
					selectFields = append(selectFields, []byte(strings.Split(currDBField, ",")[0]+",")...)
				}
			}

			// 无可查询字段
			if len(selectFields) == 0 {
				return ErrorUserInfoNoAvailiableFieldsToQuery
			}

			fieldsStr := string(selectFields)[:len(selectFields)-1]
			sqlStr := fmt.Sprintf("select %s from user where %s = ?", fieldsStr, conditionField)

			err := DB.Get(data, sqlStr, condition)
			if err == sql.ErrNoRows {
				return response.CodeUserNotExist
			} else if err != nil {
				return err
			}

			return nil
		}
	}

	return ErrorUserInfoReceiptType
}

// @Description 根据 struct tag 中 DB 查询用户信息。
// @params condition 查询条件值。
// @params data 数据接收参数 (指针类型)。
func GetUserInfoById(condition string, data any) error {
	return getUserInfo("id", condition, data)
}

// @Description 根据 struct tag 中 DB 查询用户信息。
// @params condition 查询条件值。
// @params data 数据接收参数 (指针类型)。
func GetUserInfoByUsername(condition string, data any) error {
	return getUserInfo("username", condition, data)
}

// @params username 用户名
func CheckUserExistByUsername(username string) error {
	sqlStr := "select count(id) from user where username = ?"
	var count int

	err := DB.Get(&count, sqlStr, username)

	if err != nil {
		return err
	}

	if count > 0 {
		return response.CodeUserExist
	}

	return nil
}

// @params otherFields 可选参数 [role_id, sort]
func CreateUser(user *models.User, otherFields ...interface{}) (err error) {
	sqlStr := "insert into user(uid, username, password, role_id, sort, nickname) values(?, ?, ?, ?, ?, ?)"
	randStr := helper.RandStr(15)

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
	_, err = DB.Exec(sqlStr, user.Uid, user.Username, helper.EncryptPwd(user.Password), role_id, sort, nickame)

	return
}
