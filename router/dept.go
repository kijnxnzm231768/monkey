package router

import (
	"github.com/gin-gonic/gin"
	v1 "monkey-admin/api/v1"
)

//部门接口操作
func initDeptRouter(router *gin.RouterGroup) {
	v := new(v1.DeptApi)
	group := router.Group("/dept")
	{
		//获取部门下拉树列表
		group.GET("/treeselect", v.TreeSelect)
	}
}
