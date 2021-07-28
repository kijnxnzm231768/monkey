package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/request"
)

type DeptService struct {
	deptDao dao.DeptDao
	roleDao dao.RoleDao
}

// TreeSelect 根据条件查询部门树
func (s DeptService) TreeSelect(query request.DeptQuery) *[]models.SysDept {
	treeSelect := s.deptDao.TreeSelect(query)
	return treeSelect
}

// SelectDeptListByRoleId 根据角色ID查询部门树信息
func (s DeptService) SelectDeptListByRoleId(id int64) *[]int64 {
	role := s.roleDao.SelectRoleByRoleId(id)
	return s.deptDao.SelectDeptListByRoleId(id, role.DeptCheckStrictly)
}
