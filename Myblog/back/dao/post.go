package dao

import (
	"Myblog/model"
)

// CreatePost 创建文章
func CreatePost(post *model.Post) {
	DB.Create(post)
}

// GetAllPosts 获取所有文章
func GetAllPosts() []model.Post {
	var posts []model.Post
	// 对应 SQL: SELECT * FROM posts ORDER BY created_at DESC;
	DB.Order("created_at desc").Find(&posts)
	return posts
}

// GetPostByID 根据ID获取文章
func GetPostByID(id int) model.Post {
	var post model.Post
	// Preload 是 GORM 的神器，自动把关联的 Comments 和 User 查出来
	DB.Preload("Comments.User").First(&post, id)
	return post
}
