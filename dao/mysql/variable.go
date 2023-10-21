package mysql

import "errors"

var (
	ErrorUserInfoNoAvailiableFieldsToQuery = errors.New("无效查询字段")
	ErrorUserInfoReceiptType               = errors.New("用户信息接收类型错误")
)
