package router

import (
	"github.com/gin-gonic/gin"
	v1 "monkey-admin/api/v1"
)

func initMenuRouter(router *gin.RouterGroup) {
	v := new(v1.MenuApi)
	vg := router.Group("/menu")
	{
		//加载对应角色菜单列表树
		vg.GET("/roleMenuTreeselect/:roleId", v.RoleMenuTreeSelect)
		//获取菜单下拉树列表
		vg.GET("treeselect", v.TreeSelect)
	}
}
