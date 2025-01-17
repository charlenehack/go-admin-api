package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"admin-api/common/result"
	"admin-api/common/util"
	"admin-api/pkg/db"
	"admin-api/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 用户登录
// @Summary 用户登录接口
// @Produce json
// @Description 用户登录接口
// @Param data body entity.Login true "data"
// @Success 200 {object} result.Result
// @router /api/v1/login [post]
func Login(c *gin.Context) {
	var login entity.Login
	var user entity.User
	if err := c.ShouldBindJSON(&login); err != nil {
		result.Failed(c, 408, err.Error())
		return
	}
	username := login.Username
	user = service.GetUserByUsername(username)
	// 校验用户
	if user.Username == "" {
		result.Failed(c, 419, "用户不存在")
		return
	}
	// 校验密码
	if !util.ComparePassword(login.Password, user.Password) {
		result.Failed(c, 410, "密码不正确")
		return
	}
	// 校验用户状态
	if user.Status == 0 {
		result.Failed(c, 411, "账号已停用")
		return
	}
	// 生成token
	token, _ := jwt.GenerateTokenByUser(user)
	// 左侧菜单列表
	var leftMenuVo []entity.LeftMenuVo
	leftMenuList := service.QueryLeftMenuVoList(user.ID)
	for _, value := range leftMenuList {
		menuSvoList := service.QueryMenuVoList(user.ID, value.Id)
		fmt.Println(menuSvoList)
		item := entity.LeftMenuVo{}
		item.MenuSvoList = menuSvoList
		item.Id = value.Id
		item.MenuName = value.MenuName
		item.Icon = value.Icon
		item.Url = value.Url
		leftMenuVo = append(leftMenuVo, item)
	}
	// 权限列表
	permissionList := service.QueryPermissionList(user.ID)
	var stringList = make([]string, 0)
	for _, value := range permissionList {
		stringList = append(stringList, value.Value)
	}
	result.Success(c, map[string]interface{}{"token": token, "user": user, "leftMenuList": leftMenuVo, "permissionList": stringList})
}

// 新增用户
// @Summary 新增用户接口
// @Produce json
// @Description 新增用户接口
// @Param data body entity.CreateUser true "data"
// @Success 200 {object} result.Result
// @router /api/v1/user/add [post]
// @Security ApiKeyAuth
func CreateUser(c *gin.Context) {
	var createUser entity.CreateUser
	var getUsername entity.User
	if err := c.ShouldBindJSON(&createUser); err != nil {
		result.Failed(c, 408, err.Error())
		return
	}
	getUsername = service.GetUserByUsername(createUser.Username)
	if getUsername.ID > 0 {
		result.Failed(c, 419, "用户名已存在")
		return
	}
	hashedPassword, _ := util.HashPassword(createUser.Password) // 密码加密
	user := entity.User{
		Username:    createUser.Username,
		Nickname:    createUser.Nickname,
		Password:    hashedPassword,
		Phone:       createUser.Phone,
		Email:       createUser.Email,
		Status:      createUser.Status,
		Description: createUser.Description,
	}
	tx := db.Db.Create(&user)
	if tx.RowsAffected > 0 {
		result.Success(c, true)
	} else {
		result.Failed(c, 418, "创建用户失败")
	}
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	userList, count := service.GetUserList(PageSize, PageNum, Username, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": userList})
	return
}

// 修改用户状态
func UpdateUserStatus(c *gin.Context) {
	var userStatus entity.UserStatus
	c.BindJSON(&userStatus)
	err := service.UpdateUserStatus(userStatus.ID, userStatus.Status)
	if err != nil {
		result.Failed(c, 429, "更改用户状态失败")
		return
	}
	result.Success(c, true)
}
