package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindAll(ctx context.Context) ([]*SysUser, error)
		DeleteByIds(ctx context.Context, ids string) error
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *defaultSysUserModel) FindAll(ctx context.Context) ([]*SysUser, error) {
	query := fmt.Sprintf("select %s from %s", sysUserRows, m.table)
	var resp []*SysUser
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) DeleteByIds(ctx context.Context, ids string) error {
	idsTempArr := strings.Split(ids, ",")
	query := fmt.Sprintf("delete from %s where `id` in (%s)", m.table, placeholders(len(idsTempArr)))
	var idArr []interface{}
	for _, v := range idsTempArr {
		idArr = append(idArr, v)
	}
	_, err := m.conn.ExecCtx(ctx, query, idArr...)
	return err
}

func placeholders(n int) string {
	var b strings.Builder
	for i := 0; i < n-1; i++ {
		b.WriteString("?,")
	}
	if n > 0 {
		b.WriteString("?")
	}
	return b.String()
}
