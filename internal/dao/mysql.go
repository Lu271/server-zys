package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server-zys/internal/core"
	"sync"
)

var (
	dbInstance map[string]*gorm.DB
	once       sync.Once
)

func GetDbInstance(db string) *gorm.DB {
	if db == "" {
		db = "default"
	}

	if dbInstance != nil {
		return dbInstance[db]
	}
	once.Do(func() {
		initMysql()
	})
	return dbInstance[db]
}

// initMysql 初始化全局Gorm实例
func initMysql() {
	mysqlConfig := core.GlobalConfig.Mysql
	temInstance := make(map[string]*gorm.DB)
	for _, conf := range mysqlConfig {
		db, err := gorm.Open(mysql.Open(conf.Dsn))
		if err != nil {
			return
		}
		temInstance[conf.Instance] = db
	}
	dbInstance = temInstance
}
