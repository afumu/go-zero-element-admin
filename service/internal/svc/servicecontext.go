package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zouchangfu/go-zero-element-admin/model"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewSysUserModel(conn),
	}
}
