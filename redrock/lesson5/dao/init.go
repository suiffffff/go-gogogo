package dao

//dao里面主要是对数据库的操作，分为两部分，这一部分负责链接数据库
import (
	"lesson5/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	var err error
	DB, err = gorm.Open(mysql.Open(
		"root:Paow7778200400@@tcp(127.0.0.1:3306)/xuanke?charset=utf8mb4&parseTime=True&loc=Local",
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.Student{}, &model.Course{}, &model.Teacher{}, &model.Redrockclass{}, &model.Xiaotuanti{}, &model.UserToken{})
	if err != nil {
		log.Fatal("自动迁移失败:", err)
	}
}
