package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/request"
	"monkey-admin/models/response"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userDao dao.UserDao
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
