package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRolePermissionModel = (*customSysRolePermissionModel)(nil)

type (
	// SysRolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRolePermissionModel.
	SysRolePermissionModel interface {
		sysRolePermissionModel
	}

	customSysRolePermissionModel struct {
		*defaultSysRolePermissionModel
	}
)

// NewSysRolePermissionModel returns a model for the database table.
func NewSysRolePermissionModel(conn sqlx.SqlConn) SysRolePermissionModel {
	return &customSysRolePermissionModel{
		defaultSysRolePermissionModel: newSysRolePermissionModel(conn),
	}
}
