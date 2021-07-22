package router

import (
	"github.com/gin-gonic/gin"
	"monkey-admin/pkg/filter"
	"monkey-admin/pkg/jwt"
	"monkey-admin/pkg/middleware"
	"monkey-admin/pkg/middleware/logger"
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(logger.LoggerToFile())
	router.Use(middleware.Recover)
	router.Use(jwt.JWTAuth())
	router.Use(filter.DemoHandler())
	//v1版本api
	v1Router := router.Group("/api/v1")
	{
		//登录接口
		initLoginRouter(v1Router)
		//用户路由接口
		initUserRouter(v1Router)
		//部门路由注册
		initDeptRouter(v1Router)
		//初始化字典数据路由
		initDictDataRouter(v1Router)
		//注册配置路由
		initConfigRouter(v1Router)
	}
	return router
}
