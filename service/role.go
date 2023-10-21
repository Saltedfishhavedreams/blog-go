package service

import (
	daoutils "blog/dao/daoUtils"
	"blog/models"
)

func GetRoleList() ([]*models.Role, error) {
	return daoutils.GetRoleListAndCache()
}

func GetRole(role_id string) (*models.Role, error) {
	return daoutils.GetRoleById(role_id)
}
