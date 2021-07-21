package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/pkg/cache"
)

type ConfigService struct {
	configDao dao.ConfigDao
}

// GetConfigKey 根据键名查询参数配置信息
func (s ConfigService) GetConfigKey(param string) *models.SysConfig {
	//从缓存中取出数据判断是否存在，存在直接使用，不存在就从数据库查询
	key := cache.GetRedisConfigByKey(param)
	if key != nil {
		return key
	}
	configKey := s.configDao.GetConfigKey(param)
	cache.SetRedisConfig(configKey)
	return configKey
}
