package dao

import (
	"github.com/druidcaesa/gotool"
	"github.com/go-xorm/xorm"
	"monkey-admin/models"
)

type PostDao struct {
}

func (d PostDao) selectSql(session *xorm.Session) *xorm.Session {
	return session.Table([]string{"sys_post", "p"}).
		Join("LEFT", []string{"sys_user_post", "up"}, "up.post_id = p.post_id").
		Join("LEFT", []string{"sys_user", "u"}, "u.user_id = up.user_id")
}

// SelectAll 查询所有岗位数据，数据库操作
func (d PostDao) SelectAll() []*models.SysPost {
	session := SqlDB.NewSession()
	posts := make([]*models.SysPost, 0)
	err := session.Find(&posts)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return posts
}

// SelectPostListByUserId 根据用户id查询岗位id集合
func (d PostDao) SelectPostListByUserId(userId int64) *[]int64 {
	var ids []int64
	selectSql := d.selectSql(SqlDB.NewSession())
	err := selectSql.Where("u.user_id = ?", userId).Cols("p.post_id").Find(&ids)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return &ids
}
