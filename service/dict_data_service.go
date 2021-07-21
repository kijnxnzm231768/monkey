package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/pkg/cache"
)

type DictDataService struct {
	dictDataDao dao.DictDataDao
}

// GetByType 根据字典类型查询字典数据
func (s DictDataService) GetByType(param string) []*models.SysDictData {
	//先从缓存中拉数据
	key := cache.GetRedisDictKey(param)
	if key != nil {
		return key
	} else {
		//缓存中为空，从数据库中取数据
		return s.dictDataDao.GetByType(param)
	}
}
