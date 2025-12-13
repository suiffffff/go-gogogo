package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"` // 用户名唯一
	Password string // 存放哈希后的密码，不是明文！
}
