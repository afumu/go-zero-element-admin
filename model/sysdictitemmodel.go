package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDictItemModel = (*customSysDictItemModel)(nil)

type (
	// SysDictItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictItemModel.
	SysDictItemModel interface {
		sysDictItemModel
	}

	customSysDictItemModel struct {
		*defaultSysDictItemModel
	}
)

// NewSysDictItemModel returns a model for the database table.
func NewSysDictItemModel(conn sqlx.SqlConn) SysDictItemModel {
	return &customSysDictItemModel{
		defaultSysDictItemModel: newSysDictItemModel(conn),
	}
}
