package request

import "monkey-admin/pkg/base"

// RoleQuery 角色Get请求参数
type RoleQuery struct {
	base.GlobalQuery
	RoleName string `form:"roleName"` //角色名称
	Status   string `form:"status"`   //角色状态
	RoleKey  string `form:"roleKey"`  //角色Key
}
