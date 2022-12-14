package svc

import (
	"github.com/zouchangfu/go-zero-element-admin/dao"
	"github.com/zouchangfu/go-zero-element-admin/internal/config"
	"github.com/zouchangfu/go-zero-element-admin/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	GormDB  *gorm.DB
	UserDao *dao.UserDao[model.SysUser]
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDB := dao.NewGormPlus(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		GormDB:  gormDB,
		UserDao: dao.NewUserDao[model.SysUser](),
	}
}
