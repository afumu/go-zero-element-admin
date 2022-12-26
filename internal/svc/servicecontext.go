package svc

import (
	"github.com/zouchangfu/go-zero-element-admin/internal/config"
	"github.com/zouchangfu/go-zero-element-admin/internal/dao"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	GormDB        *gorm.DB
	UserDao       *dao.UserDao[model.SysUser]
	RoleDao       *dao.RoleDao[model.SysRole]
	MenuDao       *dao.RoleDao[model.SysMenu]
	PermissionDao *dao.RoleDao[model.SysPermission]
	JwtAuth       struct {
		AccessSecret string
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDB := dao.NewGormPlus(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		GormDB:        gormDB,
		UserDao:       dao.NewUserDao[model.SysUser](),
		RoleDao:       dao.NewRoleDao[model.SysRole](),
		MenuDao:       dao.NewRoleDao[model.SysMenu](),
		PermissionDao: dao.NewRoleDao[model.SysPermission](),
	}
}
