package entity

import "time"

// 用户信息
type User struct {
	ID          int        `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	Username    string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`
	Password    string     `gorm:"column:password;varchar(64);comment:'用户密码';NOT NULL" json:"password"`
	Nickname    string     `gorm:"column:nickname;varchar(64);comment:'中文名'" json:"nickname"`
	Status      int        `gorm:"column:status;default:1;comment:'账号启用状态：1->启用，0->禁用';NOT NULL" json:"status"`
	Avatar      string     `gorm:"column:avatar;varchar(500);comment:'用户头像'" json:"avatar"`
	Email       string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`
	Phone       string     `gorm:"column:phone;varchar(64);comment:'手机号'" json:"phone"`
	Description string     `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`
	CreateAt    *time.Time `gorm:"column:createAt;comment:'创建时间'" json:"createAt"`
	UpdateAt    *time.Time `gorm:"column:updateAt;comment:'更新时间'" json:"updateAt"`
}

// 用户表名
func (User) TableName() string {
	return "sys_users"
}

// 鉴权用户结构体
type JwtUser struct {
	ID          int    `json:"ID"`          // ID
	Username    string `json:"Username"`    // 用户名
	Nickname    string `json:"Nickname"`    // 中文名
	Avatar      string `json:"Avatar"`      // 头像
	Email       string `json:"Email"`       // 邮箱
	Phone       string `json:"Phone"`       // 手机
	Description string `json:"Description"` // 描述
}

// 登录信息
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 创建用户
type CreateUser struct {
	ID          int    `json:"id"`
	Username    string `json:"username"    binding:"required"`
	Password    string `json:"password"    binding:"required"`
	Nickname    string `json:"nickname"    binding:"required"`
	Email       string `json:"email"       binding:"required,email"`
	Phone       string `json:"phone"       binding:"required,e164"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}

// 用户列表视图
type UserListVo struct {
	ID       int        `json:"id"`
	Username string     `json:"username"`
	Nickname string     `json:"nickname"`
	Status   int        `json:"status"`
	Phone    string     `json:"phone"`
	Email    string     `json:"email"`
	CreateAt *time.Time `json:"createAt" gorm:"column:createAt"`
}

// 修改用户状态请求体
type UserStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}
