package model

import (
	"time"
)

// +gplus:column=true

type SysDict struct {
	Id          int64     // 标识
	Name        string    // 字典名称
	Code        string    // 字典编码
	Description string    // 描述
	CreatedBy   int64     // 创建人
	CreatedAt   time.Time // 创建时间
	UpdatedBy   int64     // 更新人
	UpdatedAt   time.Time // 更新时间
}

func (SysDict) TableName() string {
	return "sys_dict"
}
