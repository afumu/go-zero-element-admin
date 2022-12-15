package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
)

type UserDao[T any] struct {
	gplus.CommonDao[T]
}

func NewUserDao[T any]() *UserDao[T] {
	return &UserDao[T]{}
}
