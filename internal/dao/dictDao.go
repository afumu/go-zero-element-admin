package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
)

type DictDao[T any] struct {
	gplus.CommonDao[T]
}

func NewDictDao[T any]() *DictDao[T] {
	return &DictDao[T]{}
}
