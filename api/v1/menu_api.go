package v1

import (
	"github.com/gin-gonic/gin"
	"monkey-admin/pkg/library/tree/tree_menu"
	"monkey-admin/pkg/library/user_util"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
	"strconv"
)

type MenuApi struct {
	menuService service.MenuService
}

// RoleMenuTreeSelect 加载对应角色菜单列表树
func (a MenuApi) RoleMenuTreeSelect(c *gin.Context) {
	m := make(map[string]interface{})
	param := c.Param("roleId")
	roleId, _ := strconv.ParseInt(param, 10, 64)
	//获取当前登录用户
	info := user_util.GetUserInfo(c)
	menuList := a.menuService.GetMenuTreeByUserId(info)
	menus := tree_menu.SystemMenus{}
	tree := menus.GetTree(menuList)
	ids := a.menuService.SelectMenuListByRoleId(roleId)
	m["checkedKeys"] = ids
	m["menus"] = tree
	c.JSON(200, resp.Success(m))
}

// TreeSelect 获取菜单下拉树列表
func (a MenuApi) TreeSelect(c *gin.Context) {
	info := user_util.GetUserInfo(c)
	menus := a.menuService.GetMenuTreeByUserId(info)
	systemMenus := tree_menu.SystemMenus{}
	tree := systemMenus.GetTree(menus)
	c.JSON(200, resp.Success(tree))
}
