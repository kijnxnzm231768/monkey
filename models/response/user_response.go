package response

import (
	"monkey-admin/models"
	"time"
)

// UserResponse 用户实体返回结构体
type UserResponse struct {
	UserId      int64          `json:"userId"`                    //用户ID
	DeptId      int64          `json:"deptId"`                    //部门ID
	UserName    string         `json:"userName"`                  //登录用户名
	NickName    string         `json:"nickName"`                  //用户昵称
	Email       string         `json:"email"`                     //邮箱
	PhoneNumber string         `json:"phoneNumber"`               //手机号
	Sex         string         `json:"sex"`                       //性别0男1女
	Avatar      string         `json:"avatar"`                    //头像路径
	Status      string         `json:"status"`                    //状态 0正常1停用
	DelFlag     string         `json:"delFlag"`                   //0正常1删除
	LoginIp     string         `json:"loginIp"`                   //登录ip
	LoginDate   time.Time      `json:"loginDate"`                 //登录时间
	CreateTime  time.Time      `xorm:"created" json:"createTime"` //创建时间
	CreateBy    string         `json:"createBy"`                  //创建人
	UpdateTime  time.Time      `json:"updateTime"`                //更新时间
	UpdateBy    string         `json:"updateBy"`                  //更新人
	SysDept     models.SysDept `xorm:"extends" json:"dept"`       //部门实体
}

// UserInfo 用户整体数据
type UserInfo struct {
	User    *UserResponse     `json:"user,omitempty"`    //用户数据
	Roles   []*models.SysRole `json:"roles,omitempty"`   //角色集合
	Posts   []*models.SysPost `json:"posts,omitempty"`   //部门集合
	PostIds *[]int64          `json:"postIds,omitempty"` //岗位id集合
	RoleIds *[]int64          `json:"roleIds,omitempty"` //觉得id集合
}
