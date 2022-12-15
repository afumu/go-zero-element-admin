package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
)

type RoleDao[T any] struct {
	gplus.CommonDao[T]
}

func NewRoleDao[T any]() *RoleDao[T] {
	return &RoleDao[T]{}
}
