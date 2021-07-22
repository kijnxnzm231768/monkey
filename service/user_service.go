package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/request"
	"monkey-admin/models/response"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userDao     dao.UserDao
	userPostDao dao.UserPostDao
	userRoleDao dao.UserRoleDao
}

// FindList 查询用户集合业务方法
func (s UserService) FindList(query request.UserQuery) ([]*response.UserResponse, int64) {
	return s.userDao.Find(query)
}

// GetUserById 根据id查询用户数据
func (s UserService) GetUserById(parseInt int64) *response.UserResponse {
	return s.userDao.GetUserById(parseInt)
}

// GetUserByUserName 根据用户名查询用户
func (s UserService) GetUserByUserName(name string) *models.SysUser {
	user := models.SysUser{}
	user.UserName = name
	return s.userDao.GetUserByUserName(user)
}

// CheckEmailUnique 校验邮箱是否存在
func (s UserService) CheckEmailUnique(user request.UserBody) *models.SysUser {
	return s.userDao.CheckEmailUnique(user)
}

// CheckPhoneNumUnique 校验手机号是否存在
func (s UserService) CheckPhoneNumUnique(body request.UserBody) *models.SysUser {
	return s.userDao.CheckPhoneNumUnique(body)
}

// Insert 添加用户业务逻辑
func (s UserService) Insert(body request.UserBody) bool {
	//添加用户数据库操作
	user := s.userDao.InsertUser(body)
	if user != nil {
		s.insertUserPost(user)
		s.insertUserRole(user)
		return true
	}
	return false
}

//新增用户岗位信息
func (s UserService) insertUserPost(user *request.UserBody) {
	postIds := user.PostIds
	if len(postIds) > 0 {
		sysUserPosts := make([]models.SysUserPost, 0)
		for i := 0; i < len(postIds); i++ {
			m := models.SysUserPost{
				UserId: user.UserId,
				PostId: postIds[i],
			}
			sysUserPosts = append(sysUserPosts, m)
		}
		s.userPostDao.BatchUserPost(sysUserPosts)
	}
}

//新增用户角色信息
func (s UserService) insertUserRole(user *request.UserBody) {
	roleIds := user.RoleIds
	if len(roleIds) > 0 {
		roles := make([]models.SysUserRole, 0)
		for i := 0; i < len(roleIds); i++ {
			role := models.SysUserRole{
				RoleId: roleIds[i],
				UserId: user.UserId,
			}
			roles = append(roles, role)
		}
		s.userRoleDao.BatchUserRole(roles)
	}
}

// Edit 修改用户数据
func (s UserService) Edit(body request.UserBody) int64 {
	//删除原有用户和角色关系
	s.userRoleDao.RemoveUserRole(body.UserId)
	//重新添加用具角色关系
	s.insertUserRole(&body)
	//删除原有用户岗位关系
	s.userPostDao.RemoveUserPost(body.UserId)
	//添加新的用户岗位关系
	s.insertUserPost(&body)
	//修改用户数据
	return s.userDao.Update(body)
}

// Remove 根据用户id删除用户相关数据
func (s UserService) Remove(id int64) int64 {
	//删除原有用户和角色关系
	s.userRoleDao.RemoveUserRole(id)
	//删除原有用户岗位关系
	s.userPostDao.RemoveUserPost(id)
	//删除用户数据
	return s.userDao.Remove(id)
}

// CheckUserAllowed 校验是否可以修改用户密码
func (s UserService) CheckUserAllowed(body request.UserBody) bool {
	user := models.SysUser{}
	return user.IsAdmin(body.UserId)
}

// ResetPwd 修改用户密码
func (s UserService) ResetPwd(body request.UserBody) int64 {
	return s.userDao.ResetPwd(body)
}
