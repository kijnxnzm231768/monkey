package dao

import (
	"github.com/druidcaesa/gotool"
	"monkey-admin/models"
)

type UserRoleDao struct {
}

// BatchUserRole 批量新增用户角色信息
func (d UserRoleDao) BatchUserRole(roles []models.SysUserRole) {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(models.SysUserRole{}.TableName()).Insert(&roles)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// RemoveUserRole 删除用户和角色关系
func (d UserRoleDao) RemoveUserRole(id int64) {
	role := models.SysUserRole{
		UserId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&role)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		session.Rollback()
	}
	session.Commit()
}
