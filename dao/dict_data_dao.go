package dao

import (
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"monkey-admin/models"
)

type DictDataDao struct {
}

func (d *DictDataDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_dict_data")
}

// GetByType 根据字典类型查询字典数据
func (d *DictDataDao) GetByType(param string) []*models.SysDictData {
	data := make([]*models.SysDictData, 0)
	session := d.sql(SqlDB.NewSession())
	err := session.Where("status = '0' ").And("dict_type = ?", param).OrderBy("dict_sort").Asc("dict_sort").
		Find(&data)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return data
}
