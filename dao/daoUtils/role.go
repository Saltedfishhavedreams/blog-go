package daoutils

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	"errors"
)

func GetRoleListAndCache() ([]*models.Role, error) {
	roleList, err := redis.GetRoleList()

	if err != nil {
		// 缓存错误，从mysql中重新查询，并缓存
		roleList, err := mysql.GetRoleList()

		if err != nil {
			return nil, err
		}

		err = redis.SetRoleList(roleList)
		if err != nil {
			return nil, err
		}
	} else {
		// 续期
		redis.RDB.Expire(redis.RoleList, redis.DataDefaultExpire)
	}

	return roleList, nil
}

func GetRoleById(roleId string) (*models.Role, error) {
	if roleId != "" {
		roleList, err := GetRoleListAndCache()
		if err != nil {
			return nil, err
		}

		for _, role := range roleList {
			if role.RoleId == roleId {
				return role, nil
			}
		}
	}

	return nil, errors.New("角色信息不存在")
}
