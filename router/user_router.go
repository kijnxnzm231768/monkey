package router

import (
	"github.com/gin-gonic/gin"
	v1 "monkey-admin/api/v1"
)

//用户路由
func initUserRouter(router *gin.RouterGroup) {
	userApi := new(v1.UserApi)
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", userApi.Find)
		userRouter.GET("/getInfo/:userId", userApi.GetInfo)
		userRouter.GET("/getInfo", userApi.GetInfo)
		userRouter.GET("/authRole/:userId", userApi.AuthRole)
		//新增用户
		userRouter.POST("/add", userApi.Add)
		//修改用户
		userRouter.PUT("/edit", userApi.Edit)
		//删除用户
		userRouter.DELETE("/remove/:userId", userApi.Remove)
		//重置密码
		userRouter.PUT("/resetPwd", userApi.ResetPwd)
		userRouter.GET("/export", userApi.Export)
	}
}
