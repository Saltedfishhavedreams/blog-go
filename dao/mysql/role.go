package mysql

import "blog/models"

// 查询角色列表
func GetRoleList() ([]*models.Role, error) {
	sqlStr := "select * from role"
	data := make([]*models.Role, 0)
	err := DB.Select(&data, sqlStr)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// 新增角色
func CreateRole() {}
