package router

import (
	"github.com/gin-gonic/gin"
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
	//v1版本api
	v1Router := router.Group("/api/v1")
	{
		//登录接口
		initLoginRouter(v1Router)
		//用户路由接口
		initUserRouter(v1Router)
	}
	return router
}
