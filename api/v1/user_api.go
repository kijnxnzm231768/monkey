package v1

import (
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"monkey-admin/models"
	"monkey-admin/models/request"
	"monkey-admin/models/response"
	"monkey-admin/pkg/page"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
	"strconv"
)

// UserApi 用户操作api
type UserApi struct {
	userService service.UserService
	roleService service.RoleService
	potService  service.PostService
}

// Find 查询用户列表
func (a UserApi) Find(c *gin.Context) {
	query := request.UserQuery{}
	if c.BindQuery(&query) == nil {
		list, i := a.userService.FindList(query)
		success := resp.Success(page.Page{
			Size:  query.PageSize,
			Total: i,
			List:  list,
		}, "查询成功")
		c.JSON(200, success)
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

// GetInfo 查询用户信息
func (a UserApi) GetInfo(c *gin.Context) {
	param := c.Param("userId")
	r := new(response.UserInfo)
	//查询角色
	roleAll, _ := a.roleService.SelectRoleAll(nil)
	//岗位所有数据
	postAll := a.potService.FindAll()
	//判断id传入的是否为空
	if !gotool.StrUtils.HasEmpty(param) {
		parseInt, err := strconv.ParseInt(param, 10, 64)
		if err == nil {
			//判断当前登录用户是否是admin
			m := new(models.SysUser)
			if m.IsAdmin(parseInt) {
				r.Roles = roleAll
			} else {
				roles := make([]*models.SysRole, 0)
				for _, role := range roleAll {
					if role.RoleId != 1 {
						roles = append(roles, role)
					}
				}
				r.Roles = roles
			}
			//根据id获取用户数据
			r.User = a.userService.GetUserById(parseInt)
			//根据用户ID查询岗位id集合
			r.PostIds = a.potService.SelectPostListByUserId(parseInt)
			//根据用户ID查询角色id集合
			r.RoleIds = a.roleService.SelectRoleListByUserId(parseInt)
		}
	} else {
		//id为空不取管理员角色
		roles := make([]*models.SysRole, 0)
		for _, role := range roleAll {
			if role.RoleId != 1 {
				roles = append(roles, role)
			}
		}
		r.Roles = roles
	}
	r.Posts = postAll
	c.JSON(200, resp.Success(r, "操作成功"))
}

// AuthRole 根据用户编号获取授权角色
func (a UserApi) AuthRole(c *gin.Context) {
	m := make(map[string]interface{})
	userId := c.Param("userId")
	parseInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		c.JSON(500, resp.ErrorResp(err))
	}
	user := a.userService.GetUserById(parseInt)
	//查询角色
	roles := a.roleService.GetRoleListByUserId(parseInt)
	flag := models.SysUser{}.IsAdmin(parseInt)
	if flag {
		m["roles"] = roles
	} else {
		roleList := make([]models.SysRole, 0)
		for _, role := range *roles {
			if role.RoleId != 1 {
				roleList = append(roleList, role)
			}
		}
		m["roles"] = roleList
	}
	m["user"] = user
	c.JSON(200, resp.Success(m))
}
