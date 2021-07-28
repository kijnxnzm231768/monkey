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

// SelectDeptListByRoleId 根据角色ID查询部门树信息
func (d DeptDao) SelectDeptListByRoleId(id int64, strictly bool) *[]int64 {
	list := make([]int64, 0)
	session := SqlDB.NewSession().Table([]string{"sys_dept", "d"}).Cols("d.dept_id")
	session.Join("LEFT", []string{"sys_role_dept", "rd"}, "d.dept_id = rd.dept_id").
		Where("rd.role_id = ?", id)
	if strictly {
		session.And("d.dept_id not in (select d.parent_id from sys_dept d inner join sys_role_dept rd on d.dept_id = rd.dept_id and rd.role_id = ?)", id)
	}
	err := session.OrderBy("d.parent_id").OrderBy("d.order_num").Find(&list)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &list
}
