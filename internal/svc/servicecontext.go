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
	MenuDao       *dao.MenuDao[model.SysMenu]
	PermissionDao *dao.PermissionDao[model.SysPermission]
	DictDao       *dao.DictDao[model.SysDict]
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
		MenuDao:       dao.NewMenuDao[model.SysMenu](),
		PermissionDao: dao.NewPermissionDao[model.SysPermission](),
		DictDao:       dao.NewDictDao[model.SysDict](),
	}
}
