package filter

import (
	"github.com/gin-gonic/gin"
	"monkey-admin/config"
	"net/http"
	"strings"
)

func DemoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		appServer := config.GetServerCfg()
		if appServer.DemoEnabled {
			request := inDisRequest()
			for i := 0; i < len(request); i++ {
				if strings.Contains(c.Request.RequestURI, request[i]) {
					c.JSON(http.StatusOK, gin.H{
						"status": 500,
						"msg":    "演示模式，不允许操作",
					})
					c.Abort()
					return
				}
			}

		}

	}
}

//禁用请求
func inDisRequest() []string {
	var req []string
	//一下是放行的请求
	//放行登录请求
	req = append(req, "/remove", "/edit", "/insert", "/add", "/delete", "/export", "/import")
	return req
}
