package svc

import (
	"github.com/zouchangfu/go-zero-element-admin/internal/config"
	dao2 "github.com/zouchangfu/go-zero-element-admin/internal/dao"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	GormDB  *gorm.DB
	UserDao *dao2.UserDao[model.SysUser]
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDB := dao2.NewGormPlus(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		GormDB:  gormDB,
		UserDao: dao2.NewUserDao[model.SysUser](),
	}
}
