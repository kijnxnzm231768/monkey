package v1

import (
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"monkey-admin/models"
	"monkey-admin/models/req"
	"monkey-admin/pkg/library/user_util"
	"monkey-admin/pkg/page"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
	"strconv"
	"strings"
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

// List 查询字典数据集合
func (a DictDataApi) List(c *gin.Context) {
	query := req.DiceDataQuery{}
	if c.Bind(&query) != nil {
		resp.ParamError(c)
		return
	}
	list, i := a.dictDataService.GetList(query)
	resp.OK(c, page.Page{
		List:  list,
		Total: i,
		Size:  query.PageSize,
	})
}

// Get 根据id查询字典数据
func (a DictDataApi) Get(c *gin.Context) {
	param := c.Param("dictCode")
	dictCode, _ := strconv.ParseInt(param, 10, 64)
	resp.OK(c, a.dictDataService.GetByCode(dictCode))
}

// Add 添加字典数据
func (a DictDataApi) Add(c *gin.Context) {
	data := models.SysDictData{}
	if c.Bind(&data) != nil {
		resp.ParamError(c)
		return
	}
	data.CreateBy = user_util.GetUserInfo(c).UserName
	if a.dictDataService.Insert(data) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除数据
func (a DictDataApi) Delete(c *gin.Context) {
	param := c.Param("dictCode")
	split := strings.Split(param, ",")
	dictCodeList := make([]int64, 0)
	for _, s := range split {
		diceCode, _ := strconv.ParseInt(s, 10, 64)
		dictCodeList = append(dictCodeList, diceCode)
	}
	if a.dictDataService.Remove(dictCodeList) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
