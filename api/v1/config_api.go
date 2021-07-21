package v1

import (
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
)

type ConfigApi struct {
	configService service.ConfigService
}

// GetConfigKey 根据参数键名查询参数值
func (a ConfigApi) GetConfigKey(c *gin.Context) {
	param := c.Param("configKey")
	if !gotool.StrUtils.HasEmpty(param) {
		key := a.configService.GetConfigKey(param)
		c.JSON(200, resp.Success(nil, key.ConfigValue))
	} else {
		c.JSON(500, resp.ErrorResp("参数不合法"))
	}
}
