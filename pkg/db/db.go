package db

import (
	"admin-api/common/config"
	"admin-api/pkg/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() {
	var err error
	log := log.Log()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Config.Db.Username,
		config.Config.Db.Password,
		config.Config.Db.Host,
		config.Config.Db.Port,
		config.Config.Db.Database,
		config.Config.Db.Charset)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), //  数据库操作时输出详细信息
		DisableForeignKeyConstraintWhenMigrating: true,                                // 禁用自动创建外键约束
	})
	if err != nil {
		log.Info("数据库连接失败：", err)
		panic(err)
	}
	sqlDB, err := Db.DB()
	sqlDB.SetMaxIdleConns(config.Config.Db.MaxIde)
	sqlDB.SetMaxOpenConns(config.Config.Db.MaxOpen)
}
