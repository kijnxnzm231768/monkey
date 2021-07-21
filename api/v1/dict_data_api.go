package v1

import (
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
)

type DictDataApi struct {
	dictDataService service.DictDataService
}

// GetByType 根据字典类型查询字典数据d
func (a DictDataApi) GetByType(c *gin.Context) {
	param := c.Param("dictType")
	if !gotool.StrUtils.HasEmpty(param) {
		byType := a.dictDataService.GetByType(param)
		c.JSON(200, resp.Success(byType))
	}
}
