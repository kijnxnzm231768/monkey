package models

import (
	"time"
)

// SysDictType 字典类型表
type SysDictType struct {
	DictId     int64     `xorm:"pk autoincr" json:"dictId"`    //字典ID
	DictName   string    `xorm:"varchar(128)" json:"dictName"` //字典名称
	DictType   string    `xorm:"varchar(128)" json:"dictType"` //字典类型
	Status     string    `xorm:"char(1)" json:"status"`        //状态 0正常1停用
	Remark     string    `xorm:"varchar(512)" json:"remark"`   //备注
	CreateTime time.Time `xorm:"created" json:"createTime"`    //创建时间
	CreateBy   string    `json:"createBy"`                     //创建人
	UpdateTime time.Time `json:"updateTime"`                   //更新时间
	UpdateBy   string    `json:"updateBy"`                     //更新人
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
