package models

import "time"

type SysRole struct {
	RoleId            int64     `xorm:"pk autoincr" json:"roleId"`   //角色id
	RoleName          string    `xorm:"varchar(64)" json:"roleName"` //角色名称
	RoleKey           string    `xorm:"varchar(64)" json:"roleKey"`  //角色权限标识
	RoleSort          int       `xorm:"int" json:"roleSort"`         //角色顺序
	DataScope         string    `json:"dataScope"`                   //数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly bool      `json:"menuCheckStrictly"`           //菜单树选择项是否关联显示
	DeptCheckStrictly bool      `json:"deptCheckStrictly"`           //部门树选择项是否关联显示
	Status            string    `xorm:"char(1)" json:"status"`       //角色状态 0正常1停用
	DelFlag           string    `xorm:"char(1)" json:"delFlag"`      //删除标记0正常1删除
	CreateTime        time.Time `xorm:"created" json:"createTime"`   //创建时间
	CreateBy          string    `json:"createBy"`                    //创建人
	UpdateTime        time.Time `json:"updateTime"`                  //更新时间
	UpdateBy          string    `json:"updateBy"`                    //更新人
	Remark            string    `json:"remark"`                      //备注
	MenuIds           []int64   `xorm:"-" json:"menuIds"`            //菜单组
	DeptIds           []int64   `xorm:"-" json:"deptIds"`            //部门组
}

func (r SysRole) TableName() string {
	return "sys_role"
}
