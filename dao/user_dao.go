package dao

import (
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"monkey-admin/models"
	"monkey-admin/models/request"
	"monkey-admin/models/response"
	"monkey-admin/pkg/page"
)

type UserDao struct {
}

//查询公共sql
func (d UserDao) querySql() *xorm.Session {
	return SqlDB.NewSession().Table([]string{"sys_user", "u"}).
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id")
}

// Find 查询用户集合
func (d UserDao) Find(query request.UserQuery) ([]*response.UserResponse, int64) {
	resp := make([]*response.UserResponse, 0)
	sql := d.querySql()
	if !gotool.StrUtils.HasEmpty(query.UserName) {
		sql.And("u.user_name like concat('%',?,'%')", query.UserName)
	}
	if !gotool.StrUtils.HasEmpty(query.Status) {
		sql.And("i.status = ?", query.Status)
	}
	if !gotool.StrUtils.HasEmpty(query.PhoneNumber) {
		sql.And("u.phone_number like concat('%',?,'%')", query.PhoneNumber)
	}
	if !gotool.StrUtils.HasEmpty(query.BeginTime) {
		sql.And("date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')", query.BeginTime)
	}
	if !gotool.StrUtils.HasEmpty(query.EndTime) {
		sql.And("date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')", query.EndTime)
	}
	if query.DeptId > 0 {
		sql.And("u.dept_id = ? OR u.dept_id in ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(?, ancestors))", query.DeptId, query.DeptId)
	}
	total, _ := page.GetTotal(sql.Clone())
	err := sql.Limit(query.PageSize, page.StartSize(query.PageNum, query.PageSize)).Find(&resp)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil, 0
	}
	return resp, total
}

// GetUserById 根据id查询用户数据
func (d UserDao) GetUserById(userId int64) *response.UserResponse {
	var resp response.UserResponse
	_, err := d.querySql().Where("u.user_id = ?", userId).Get(&resp)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &resp
}

// GetUserByUserName 根据用户名查询用户数据
func (d UserDao) GetUserByUserName(user models.SysUser) *models.SysUser {
	_, err := SqlDB.Get(&user)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &user
}
