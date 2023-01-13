package utils

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"time"
)

func CovertTimeToString() copier.Option {
	return copier.Option{
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					if s.UnixMilli() < 0 {
						return "", nil
					}
					return s.Format("2006-01-02 15:04:05"), nil
				},
			},
		},
	}
}
