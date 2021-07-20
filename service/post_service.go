package service

import (
	"monkey-admin/dao"
	"monkey-admin/models"
)

type PostService struct {
	postDao dao.PostDao
}

// FindAll 查询所有岗位业务方法
func (s PostService) FindAll() []*models.SysPost {
	return s.postDao.SelectAll()
}

// SelectPostListByUserId 根据用户id查询岗位id集合
func (s PostService) SelectPostListByUserId(userId int64) *[]int64 {
	return s.postDao.SelectPostListByUserId(userId)
}
