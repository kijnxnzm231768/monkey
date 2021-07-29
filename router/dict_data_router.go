package router

import (
	"github.com/gin-gonic/gin"
	v1 "monkey-admin/api/v1"
)

//初始化字典数据路由
func initDictDataRouter(router *gin.RouterGroup) {
	v := new(v1.DictDataApi)
	group := router.Group("/dict/data")
	{
		//根据字典类型查询字典数据信息
		group.GET("/type/:dictType", v.GetByType)
	}
}
