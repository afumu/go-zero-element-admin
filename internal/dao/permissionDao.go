package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
)

type PermissionDao[T any] struct {
	gplus.CommonDao[T]
}

func NewPermissionDao[T any]() *PermissionDao[T] {
	return &PermissionDao[T]{}
}
