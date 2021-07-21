package dao

import (
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"monkey-admin/models"
)

type ConfigDao struct {
}

func (d ConfigDao) sql(session *xorm.Session) *xorm.Session {
	return session.Table("sys_config")
}

// GetConfigKey 根据键名查询参数配置信息
func (d ConfigDao) GetConfigKey(param string) *models.SysConfig {
	config := models.SysConfig{}
	_, err := d.sql(SqlDB.NewSession()).Where("config_key = ?", param).Get(&config)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &config
}
