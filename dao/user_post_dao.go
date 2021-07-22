package dao

import (
	"github.com/druidcaesa/gotool"
	"monkey-admin/models"
)

type UserPostDao struct {
}

// BatchUserPost 批量新增用户岗位信息
func (d UserPostDao) BatchUserPost(posts []models.SysUserPost) {
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Table(models.SysUserPost{}.TableName()).Insert(&posts)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		session.Rollback()
		return
	}
	session.Commit()
}

// RemoveUserPost 删除用户和岗位关系
func (d UserPostDao) RemoveUserPost(id int64) {
	post := models.SysUserPost{
		UserId: id,
	}
	session := SqlDB.NewSession()
	session.Begin()
	_, err := session.Delete(&post)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		session.Rollback()
	}
	session.Commit()
}
