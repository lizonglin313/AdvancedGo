package models

import (
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

// 对数据库进行初始化

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		// dbType string
		dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "fail to get section 'database': %v", err)
	}

	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//  db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	user,
	//	password,
	//	host,
	//	dbName))
	// gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 通过设置表前缀更改默认表名
			SingularTable: true,        // 禁用表名复数形式
		},
		Logger: logger.Default.LogMode(logger.Info), // 启用日志 使用日志级别为 logger
	})
	if err != nil {
		log.Fatalf("error of connection to db: %v", err)
	}

	// 获取通用数据对象 sql.DB 使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error of get db.DB(): %v", err)
	}
	sqlDB.SetMaxIdleConns(10)  // 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100) // 设置打开数据库连接的最大数量

}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error of get db.DB(): %v", err)
	}
	defer sqlDB.Close()
}
