package models

import "time"

// SysDept 部门结构体
type SysDept struct {
	DeptId     int64     `xorm:"pk autoincr" json:"deptId"`
	Ancestors  string    `xorm:"varchar(50)" json:"ancestors"`
	DeptName   string    `xorm:"varchar(128)" json:"deptName"`
	OrderNum   int       `json:"orderNum"`
	Leader     string    `xorm:"varchar(20)" json:"leader"`
	ParentId   int64     `json:"parentId"`
	Phone      string    `xorm:"varchar(11)" json:"phone"`
	Status     string    `xorm:"char(1)" json:"status"`
	DelFlag    string    `xorm:"char(1)" json:"delFlag"`
	CreateTime time.Time `xorm:"created" json:"createTime"` //创建时间
	CreateBy   string    `json:"createBy"`                  //创建人
	UpdateTime time.Time `json:"updateTime"`                //更新时间
	UpdateBy   string    `json:"updateBy"`                  //更新人
}

func (SysDept) TableName() string {
	return "sys_dept"
}
