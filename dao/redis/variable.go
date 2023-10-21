package redis

import "errors"

/**
redis 存储数据前缀
*/
const (
	RedisRolePrefix = "ROLE-"
	RedisMenuPrefix = "Menu-"
)

/**
redis 错误提示
*/
var (
	ErrorRoleIsNotExist = errors.New("角色信息不存在")
)
