package model

import (
	"time"
)

// +gplus:column=true

type SysPermission struct {
	Id        int64     // 标识
	menuId    string    // 菜单标识
	Name      string    // 权限名称
	Code      string    // 权限编码
	Status    string    // 权限状态
	CreatedBy int64     // 创建人
	CreatedAt time.Time // 创建时间
	UpdatedBy int64     // 更新人
	UpdatedAt time.Time // 更新时间
}

func (SysPermission) TableName() string {
	return "sys_permission"
}
