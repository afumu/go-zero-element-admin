package model

import (
	"time"
)

// +gplus:column=true

type SysMenu struct {
	Id          int64     // 标识
	ParentId    string    // 父标识
	Name        string    // 菜单名称
	Url         string    // 菜单路径
	Sort        int       // 菜单排序
	Icon        string    // 菜单图标
	Description string    // 描述
	CreatedBy   int64     // 创建人
	CreatedAt   time.Time // 创建时间
	UpdatedBy   int64     // 更新人
	UpdatedAt   time.Time // 更新时间
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
