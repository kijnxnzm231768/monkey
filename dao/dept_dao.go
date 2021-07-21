package dao

import (
	"github.com/druidcaesa/gotool"
	"monkey-admin/models"
	"monkey-admin/models/request"
)

type DeptDao struct {
}

// TreeSelect 根据条件查询部门集合
func (d DeptDao) TreeSelect(query request.DeptQuery) *[]models.SysDept {
	depts := make([]models.SysDept, 0)
	session := SqlDB.NewSession().Where("del_flag = '0'")
	if query.ParentId > 0 {
		session.And("parent_id = ?", query.ParentId)
	}
	if !gotool.StrUtils.HasEmpty(query.DeptName) {
		session.And("dept_name like concat('%', ?, '%')", query.DeptName)
	}
	if !gotool.StrUtils.HasEmpty(query.Status) {
		session.And("status = ?", query.Status)
	}
	err := session.OrderBy("parent_id").OrderBy("order_num").Find(&depts)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &depts
}
