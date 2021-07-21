package v1

import (
	"github.com/gin-gonic/gin"
	"monkey-admin/models/request"
	"monkey-admin/pkg/library/tree/tree_dept"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
)

type DeptApi struct {
	deptService service.DeptService
}

// TreeSelect 查询部门菜单树
func (a DeptApi) TreeSelect(c *gin.Context) {
	query := request.DeptQuery{}
	if c.BindQuery(&query) == nil {
		treeSelect := a.deptService.TreeSelect(query)
		list := tree_dept.DeptList{}
		list = *treeSelect
		array := list.ConvertToINodeArray(treeSelect)
		tree := tree_dept.GenerateTree(array, nil)
		c.JSON(200, resp.Success(tree))
	} else {
		c.JSON(500, resp.ErrorResp("参数绑定错误"))
	}
}
