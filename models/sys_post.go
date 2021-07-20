package models

import "time"

// SysPost 岗位表对应结构体
type SysPost struct {
	PostId     int64     `xorm:"pk autoincr" json:"postId"`   //岗位ID
	PostCode   string    `xorm:"varchar(64)" json:"postCode"` //岗位编码
	PostName   string    `xorm:"varchar(64)" json:"postName"` //岗位名称
	Status     string    `xorm:"char(1)" json:"status"`       //状态 0正常 1停用
	Remark     string    `xorm:"varchar(512)" json:"remark"`  //备注
	CreateTime time.Time `xorm:"created" json:"createTime"`   //创建时间
	CreateBy   string    `json:"createBy"`                    //创建人
	UpdateTime time.Time `json:"updateTime"`                  //更新时间
	UpdateBy   string    `json:"updateBy"`                    //更新人
}

func (SysPost) TableName() string {
	return "sys_post"
}
