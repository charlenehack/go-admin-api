package service

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"fmt"
)

// 根据用户名查询用户信息
func GetUserByUsername(username string) (user entity.User) {
	db.Db.Where("username = ?", username).First(&user)
	return user
}

// 修改用户状态
func UpdateUserStatus(id, status int) (err error) {
	var user entity.User
	res := db.Db.Model(&user).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("未查询到用户")
	}
	return nil
}

// 查询用户列表
func GetUserList(PageSize, PageNum int, Username, Status, BeginTime, EndTime string) (sysUserVo []entity.UserListVo, count int64) {
	curDb := db.Db.Table("sys_users")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username) // 根据用户名进行查询
	}
	if Status != "" {
		curDb = curDb.Where("status = ?", Status) // 根据用户状态进行查询
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("createAt BETWEEN ? AND ?", BeginTime, EndTime) // 根据创建时间进行查询
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("createAt DESC").Find(&sysUserVo)

	return sysUserVo, count
}

// 根据用户ID获取二级菜单列表
func QueryMenuVoList(UserId, MenuId int) (MenuVo []entity.MenuSvo) {
	const status, menuStatus, menuType = 1, 1, 2
	db.Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		//		Joins("LEFT JOIN sys_users su ON su.id = sar.admin_id").
		//		Where("sr.status = ?", status).
		//		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		//		Where("su.id = ?", Id).
		Order("sm.sort").
		Scan(&MenuVo)
	return MenuVo
}

// 根据用户ID获取左侧一级菜单列表
func QueryLeftMenuVoList(Id int) (leftMenuVo []entity.LeftMenuVo) {
	const status, menuStatus, menuType = 1, 1, 1
	db.Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		//		Joins("LEFT JOIN sys_users su ON su.id = sar.admin_id").
		//		Where("sr.status = ?", status).
		//		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		//		Where("su.id = ?", Id).
		Order("sm.sort").
		Scan(&leftMenuVo)
	return leftMenuVo
}

// 根据用户ID获取权限列表
func QueryPermissionList(Id int) (valueVo []entity.PerssionVo) {
	const status, menuStatus, menuType = 1, 1, 3
	db.Db.Table("sys_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		//		Joins("LEFT JOIN sys_users su ON su.id = sar.admin_id").
		//		Where("sr.status = ?", status).
		//		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		//		Where("su.id = ?", Id).
		Order("sm.sort").
		Scan(&valueVo)
	return valueVo
}
