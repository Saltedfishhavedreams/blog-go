package redis

import (
	"blog/models"
	"encoding/json"
)

const (
	RoleList = RedisRolePrefix + "ROLE-LIST"
)

func SetRoleList(roleList []*models.Role) error {
	bytes, err := json.Marshal(roleList)
	if err != nil {
		return err
	}

	return RDB.Set(RoleList, string(bytes), DataDefaultExpire).Err()
}

func GetRoleList() ([]*models.Role, error) {
	roleList := make([]*models.Role, 0)
	jsonStr, err := RDB.Get(RoleList).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonStr), &roleList)
	if err != nil {
		return nil, err
	}
	return roleList, nil
}
