package model

import (
	"time"
)

// +gplus:column=true

type SysRole struct {
	Id          int64     // 标识
	Name        string    // 角色名
	Code        string    // 角色编码
	Description string    // 描述
	CreatedBy   int64     // 创建人
	CreatedAt   time.Time // 创建时间
	UpdatedBy   int64     // 更新人
	UpdatedAt   time.Time // 更新时间
}

func (SysRole) TableName() string {
	return "sys_role"
}
