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
	}
}
