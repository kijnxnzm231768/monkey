package request

import "monkey-admin/pkg/base"

// UserQuery 用户get请求数据参数
type UserQuery struct {
	base.GlobalQuery
	UserName    string `form:"userName"`    //用户名
	Status      string `form:"status"`      //状态
	PhoneNumber string `form:"phoneNumber"` //手机号
	DeptId      int64  `form:"deptId"`      //部门id
}
