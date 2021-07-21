package dao

import (
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"monkey-admin/models"
	"monkey-admin/models/request"
	"monkey-admin/pkg/page"
)

type RoleDao struct {
}

//角色公用sql
func (d RoleDao) sqlSelectJoin() *xorm.Session {
	return SqlDB.Table([]string{models.SysRole{}.TableName(), "r"}).
		Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id").
		Join("LEFT", []string{"sys_dept", "d"}, "u.dept_id = d.dept_id")
}

//用户角色关系查询sql
func (d RoleDao) sqlSelectRoleAndUser() *xorm.Session {
	return SqlDB.Table([]string{models.SysRole{}.TableName(), "r"}).
		Join("LEFT", []string{"sys_user_role", "ur"}, "ur.role_id = r.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = ur.user_id")
}

// SelectRoleList 根据条件查询角色数据
func (d RoleDao) SelectRoleList(q *request.RoleQuery) ([]*models.SysRole, int64) {
	roles := make([]*models.SysRole, 0)
	session := d.sqlSelectJoin()
	if !gotool.StrUtils.HasEmpty(q.RoleName) {
		session.And("r.role_name like concat('%', ?, '%')", q.RoleName)
	}
	if !gotool.StrUtils.HasEmpty(q.Status) {
		session.And("r.status = ?", q.Status)
	}
	if !gotool.StrUtils.HasEmpty(q.RoleKey) {
		session.And("r.role_key like concat('%', ?, '%')", q.RoleKey)
	}
	if !gotool.StrUtils.HasEmpty(q.BeginTime) {
		timestamp, _ := gotool.DateUtil.InterpretStringToTimestamp(q.BeginTime, "YYYY-MM-DD  hh:mm:ss")
		session.And("date_format(r.create_time,'%y%m%d') &gt;= date_format(?,'%y%m%d')", timestamp)
	}
	if !gotool.StrUtils.HasEmpty(q.EndTime) {
		timestamp, _ := gotool.DateUtil.InterpretStringToTimestamp(q.EndTime, "YYYY-MM-DD  hh:mm:ss")
		session.And("date_format(r.create_time,'%y%m%d') &lt;= date_format(?,'%y%m%d')", timestamp)
	}
	total, _ := page.GetTotal(session.Clone())
	err := session.Limit(q.PageSize, page.StartSize(q.PageNum, q.PageSize)).Find(&roles)
	if err != nil {
		return nil, 0
	}
	return roles, total
}

// SelectRoleAll 查询所有角色
func (d RoleDao) SelectRoleAll() []*models.SysRole {
	sql := d.sqlSelectJoin()
	roles := make([]*models.SysRole, 0)
	err := sql.Find(&roles)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return roles
}

// SelectRoleListByUserId 根据用户id查询用户角色id集合
func (d RoleDao) SelectRoleListByUserId(userId int64) *[]int64 {
	sqlSelectRoleAndUser := d.sqlSelectRoleAndUser()
	var roleIds []int64
	err := sqlSelectRoleAndUser.Cols("r.role_id").Where("u.user_id = ?", userId).Find(&roleIds)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &roleIds
}

// GetRolePermissionByUserId 查询用户角色集合
func (d RoleDao) GetRolePermissionByUserId(id int64) *[]string {
	var roleKeys []string
	err := d.sqlSelectJoin().Cols("r.role_key").Where("r.del_flag = '0'").And("ur.user_id = ?", id).Find(&roleKeys)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &roleKeys
}
