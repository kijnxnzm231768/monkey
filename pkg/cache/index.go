package cache

import (
	"github.com/druidcaesa/gotool"
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/pkg/constant"
)

// GetRedisDictKey 根据key获取缓存中的字典数据
func GetRedisDictKey(key string) []*models.SysDictData {
	get, err := dao.RedisDB.GET(key)
	if err != nil {
		gotool.Logs.ErrorLog().Fatalf(constant.RedisConstant{}.GetRedisError(), err.Error())
		return nil
	}
	list := models.SysDictData{}.UnmarshalDictList(get)
	return list
}

// SetRedisDict 保存字典数据
func SetRedisDict(list []*models.SysDictData) {
	dictList := models.SysDictData{}.MarshalDictList(list)
	dao.RedisDB.SET(list[0].DictType, dictList)
}

// GetRedisConfigByKey 根据key从缓存中获取配置数据
func GetRedisConfigByKey(key string) *models.SysConfig {
	get, err := dao.RedisDB.GET(key)
	if err != nil {
		gotool.Logs.ErrorLog().Fatalf(constant.RedisConstant{}.GetRedisError(), err.Error())
		return nil
	}
	obj := models.SysConfig{}.UnmarshalDictObj(get)
	return obj
}

// SetRedisConfig 将配置存入缓存
func SetRedisConfig(config *models.SysConfig) {
	s := models.SysConfig{}.MarshalDictObj(*config)
	dao.RedisDB.SET(config.ConfigKey, s)
}
