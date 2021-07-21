package router

import (
	"github.com/gin-gonic/gin"
	v1 "monkey-admin/api/v1"
)

func initConfigRouter(router *gin.RouterGroup) {
	v := new(v1.ConfigApi)
	group := router.Group("/config")
	{
		//根据参数键名查询参数值
		group.GET("/configKey/:configKey", v.GetConfigKey)
	}
}
