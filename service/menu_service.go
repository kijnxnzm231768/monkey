package service

import (
	"github.com/gogf/gf/util/gconv"
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/response"
	"monkey-admin/pkg/library/utils"
)

type MenuService struct {
	menuDao dao.MenuDao
}

// GetMenuTreeByUserId 根据用户ID查询菜单
func (s MenuService) GetMenuTreeByUserId(user *response.UserResponse) []map[string]interface{} {
	var menuList *[]models.SysMenu
	//判断是否是管理员
	flag := models.SysUser{}.IsAdmin(user.UserId)
	if flag {
		menuList = s.menuDao.GetMenuAll()
	} else {
		menuList = s.menuDao.GetMenuByUserId(user.UserId)
	}
	list := gconv.SliceMap(menuList)
	list = utils.PushSonToParent(list, 0, "parentId", "menuId", "children", "", nil, true)
	return list
}
