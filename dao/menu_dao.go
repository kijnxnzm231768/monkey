package dao

import (
	"github.com/druidcaesa/gotool"
	"monkey-admin/models"
)

type MenuDao struct {
}

// GetMenuPermission 根据用户ID查询权限
func (d MenuDao) GetMenuPermission(id int64) *[]string {
	var perms []string
	session := SqlDB.Table([]string{"sys_menu", "m"})
	err := session.Distinct("m.perms").
		Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
		Join("LEFT", []string{"sys_role", "r"}, "r.role_id = ur.role_id").
		Where("m.status = '0'").And("r.status = '0'").And("ur.user_id = ?", id).Find(&perms)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &perms
}

// GetMenuAll 查询所有菜单数据
func (d MenuDao) GetMenuAll() *[]models.SysMenu {
	menus := make([]models.SysMenu, 0)
	session := SqlDB.Table([]string{models.SysMenu{}.TableName(), "m"})
	err := session.Distinct("m.menu_id").Cols("m.parent_id", "m.menu_name", "m.path", "m.component", "m.visible", "m.status", "m.perms", "m.is_frame", "m.is_cache", "m.menu_type", "m.icon", "m.order_num", "m.create_time").
		Where("m.menu_type in ('M', 'C')").And("m.status = 0").OrderBy("m.parent_id").OrderBy("m.order_num").Find(&menus)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &menus
}

// GetMenuByUserId 根据用户ID查询菜单
func (d MenuDao) GetMenuByUserId(id int64) *[]models.SysMenu {
	menus := make([]models.SysMenu, 0)
	session := SqlDB.Table([]string{models.SysMenu{}.TableName(), "m"})
	err := session.Distinct("m.menu_id").Cols("m.parent_id", "m.menu_name", "m.path", "m.component", "m.visible", "m.status", "m.perms", "m.is_frame", "m.is_cache", "m.menu_type", "m.icon", "m.order_num", "m.create_time").
		Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id").
		Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id").
		Join("LEFT", []string{"sys_role", "ro"}, "ur.role_id = ro.role_id").
		Join("LEFT", []string{"sys_user", "u"}, "ur.user_id = u.user_id").Where("u.user_id = ?", id).
		And("m.menu_type in ('M', 'C')").And("m.status = 0").OrderBy("m.parent_id").OrderBy("m.order_num").Find(&menus)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &menus
}

// GetMenuByRoleId 根据角色ID查询菜单树信息
func (d MenuDao) GetMenuByRoleId(id int64, strictly bool) *[]int64 {
	list := make([]int64, 0)
	session := SqlDB.NewSession().Table([]string{"sys_menu", "m"})
	session.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id")
	session.Where("rm.role_id = ?", id)
	if strictly {
		session.And("m.menu_id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.menu_id = rm.menu_id and rm.role_id = ?)", id)
	}
	err := session.OrderBy("m.parent_id").OrderBy("m.order_num").Cols("m.menu_id").Find(&list)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &list
}
