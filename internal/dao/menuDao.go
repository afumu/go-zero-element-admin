package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/zouchangfu/go-zero-element-admin/internal/model"
	"gorm.io/gorm"
)

type MenuDao[T any] struct {
	gplus.CommonDao[T]
}

func NewMenuDao[T any]() *MenuDao[T] {
	return &MenuDao[T]{}
}

func (menuDao MenuDao[T]) GetUserMenu() ([]*model.SysMenu, *gorm.DB) {
	var menus []*model.SysMenu
	sqlResult := menuDao.Db().Table("sys_menu as m").
		Joins("left join sys_role_menu rm on  rm.menu_id = m.id").
		Joins("left join sys_role r on r.id =  rm.role_id").
		Joins("left join sys_user_role ur on ur.user_id = 1").Select("m.*").Scan(&menus)
	return menus, sqlResult
}
