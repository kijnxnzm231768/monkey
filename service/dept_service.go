package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/request"
)

type DeptService struct {
	deptDao dao.DeptDao
}

// TreeSelect 根据条件查询部门树
func (s DeptService) TreeSelect(query request.DeptQuery) *[]models.SysDept {
	treeSelect := s.deptDao.TreeSelect(query)
	return treeSelect
}
