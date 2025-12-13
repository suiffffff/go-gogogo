package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserID  uint // 关联用户ID
	User    User // 关联用户对象
	PostID  uint // 关联文章ID
}
