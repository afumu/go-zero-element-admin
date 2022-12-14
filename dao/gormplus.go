package dao

import (
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewGormPlus(datasource string) *gorm.DB {
	gormDb, err := gorm.Open(mysql.Open(datasource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	gplus.Init(gormDb)
	return gormDb
}
