package datasource

import (
	"iris-demo/models"
)

// Createtable 初始化表 如果不存在该表 则自动创建
func Createtable() {
	err := GetDB().AutoMigrate(
		&models.User{},
	)
	if err != nil {
		return
	}
}