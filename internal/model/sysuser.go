package model

import (
	"database/sql"
	"time"
)

// +gplus:column=true

type SysUser struct {
	Id        int64        // 标识
	Username  string       // 用户账号
	Realname  string       // 真实姓名
	Password  string       // 密码
	Salt      string       // md5密码盐
	Sex       int64        // 性别【0:未知，1-男，2-女】
	Avatar    string       // 头像
	Status    int64        // 状态【0-正常,1-冻结】
	DeletedAt sql.NullTime // 删除状态【0-正常，1-已删除】
	CreatedBy int64        // 创建人
	CreatedAt time.Time    // 创建时间
	UpdatedBy int64        // 更新人
	UpdatedAt time.Time    // 更新时间
}

func (SysUser) TableName() string {
	return "sys_user"
}
