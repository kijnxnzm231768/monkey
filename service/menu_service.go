package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/response"
)

type MenuService struct {
	menuDao dao.MenuDao
	roleDao dao.RoleDao
}

// GetMenuTreeByUserId 根据用户ID查询菜单
func (s MenuService) GetMenuTreeByUserId(user *response.UserResponse) *[]models.SysMenu {
	var menuList *[]models.SysMenu
	//判断是否是管理员
	flag := models.SysUser{}.IsAdmin(user.UserId)
	if flag {
		menuList = s.menuDao.GetMenuAll()
	} else {
		menuList = s.menuDao.GetMenuByUserId(user.UserId)
	}
	return menuList
}

// SelectMenuListByRoleId 根据角色ID查询菜单树信息
func (s MenuService) SelectMenuListByRoleId(id int64) *[]int64 {
	role := s.roleDao.SelectRoleByRoleId(id)
	return s.menuDao.GetMenuByRoleId(id, role.MenuCheckStrictly)
}
