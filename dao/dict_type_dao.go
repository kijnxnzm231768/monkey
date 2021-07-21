package dao

import (
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"monkey-admin/models"
)

type DictTypeDao struct {
}

func (d DictTypeDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_dict_type")
}

// FindAll 查询所有字典类型数据
func (d DictTypeDao) FindAll() []*models.SysDictType {
	types := make([]*models.SysDictType, 0)
	err := d.sql(SqlDB.NewSession()).Where("status = '0'").Find(&types)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return types
}
