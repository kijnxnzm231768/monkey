package v1

import (
	"github.com/gin-gonic/gin"
	"monkey-admin/models/request"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"
)

type LoginApi struct {
	loginService      service.LoginService
	roleService       service.RoleService
	permissionService service.PermissionService
	menuService       service.MenuService
}

// Login 登录
func (a LoginApi) Login(c *gin.Context) {
	loginBody := request.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		m := make(map[string]string)
		login, s := a.loginService.Login(loginBody.UserName, loginBody.Password)
		if login {
			m["token"] = s
			c.JSON(200, resp.Success(m))
		} else {
			c.JSON(401, resp.ErrorResp(s))
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}

// GetUserInfo 获取用户信息
func (a LoginApi) GetUserInfo(c *gin.Context) {
	m := make(map[string]interface{})
	user := a.loginService.LoginUser(c)
	//查询用户角色集合
	roleKeys := a.permissionService.GetRolePermissionByUserId(user)
	// 权限集合
	perms := a.permissionService.GetMenuPermission(user)
	m["roles"] = roleKeys
	m["permissions"] = perms
	m["user"] = user
	c.JSON(200, resp.Success(m))
}

// GetRouters 根据用户ID查询菜单
func (a LoginApi) GetRouters(c *gin.Context) {
	//获取等钱登录用户
	user := a.loginService.LoginUser(c)
	menus := a.menuService.GetMenuTreeByUserId(user)
	c.JSON(200, resp.Success(menus))
}
