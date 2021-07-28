package v1

import (
	"github.com/gin-gonic/gin"
	"monkey-admin/models/request"
	"monkey-admin/pkg/library/tree/tree_dept"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
	"strconv"
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
		c.JSON(200, resp.Success(list.GetTree(treeSelect)))
	} else {
		c.JSON(500, resp.ErrorResp("参数绑定错误"))
	}
}

// RoleDeptTreeSelect 加载对应角色部门列表树
func (a DeptApi) RoleDeptTreeSelect(c *gin.Context) {
	m := make(map[string]interface{})
	param := c.Param("roleId")
	roleId, _ := strconv.ParseInt(param, 10, 64)
	checkedKeys := a.deptService.SelectDeptListByRoleId(roleId)
	m["checkedKeys"] = checkedKeys
	treeSelect := a.deptService.TreeSelect(request.DeptQuery{})
	list := tree_dept.DeptList{}
	tree := list.GetTree(treeSelect)
	m["depts"] = tree
	resp.OK(c, m)
}
