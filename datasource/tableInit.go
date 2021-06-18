package datasource

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iris-demo/conf"
	"strings"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	// root:root123@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local
	path := strings.Join([]string{conf.Sysconfig.DBUserName, ":", conf.Sysconfig.DBPassword, "@(", conf.Sysconfig.DBIp, ":", conf.Sysconfig.DBPort, ")/", conf.Sysconfig.DBName, "?charset=utf8&parseTime=true"}, "")
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: path,
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err !=nil{
		panic(err)
	}

	Createtable()
}