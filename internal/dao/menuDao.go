package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
)

type MenuDao[T any] struct {
	gplus.CommonDao[T]
}

func NewMenuDao[T any]() *MenuDao[T] {
	return &MenuDao[T]{}
}
