package dao

import (
	"Myblog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	// dsn 格式: 用户名:密码@tcp(IP:端口)/数据库名?配置
	// 如果是本地测试用 127.0.0.1，部署时如果是本机也用 127.0.0.1
	dsn := "root:mynameocy@@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}

	// 自动迁移所有模型
	err = DB.AutoMigrate(&model.Post{}, &model.User{}, &model.Comment{})
	if err != nil {
		log.Fatalf("建表失败: %v", err)
	}
}
