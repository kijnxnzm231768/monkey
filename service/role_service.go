package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/request"
)

type RoleService struct {
	roleDao dao.RoleDao
}

// SelectRoleAll 查询所有角色
func (s RoleService) SelectRoleAll(query *request.RoleQuery) ([]*models.SysRole, int64) {
	if query == nil {
		all := s.roleDao.SelectRoleAll()
		return all, 0
	}
	return s.roleDao.SelectRoleList(query)
}

// SelectRoleListByUserId 根据用户id查询角色id集合
func (s RoleService) SelectRoleListByUserId(parseInt int64) *[]int64 {
	return s.roleDao.SelectRoleListByUserId(parseInt)
}
