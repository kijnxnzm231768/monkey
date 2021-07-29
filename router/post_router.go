package router

import (
	"github.com/gin-gonic/gin"
	v1 "monkey-admin/api/v1"
)

//初始化岗位路由
func initPostRouter(router *gin.RouterGroup) {
	v := new(v1.PostApi)
	group := router.Group("/post")
	{
		//查询岗位数据
		group.GET("/list", v.List)
		//添加岗位
		group.POST("/add", v.Add)
		//查询岗位详情
		group.GET("/:postId", v.Get)
		//删除岗位数据
		group.DELETE("/:postId", v.Delete)
	}
}
