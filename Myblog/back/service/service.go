package service

import (
	"Myblog/dao"
	"Myblog/model"
)

type PostService struct{}

// GetPostList 获取文章列表
func (s *PostService) GetPostList() []model.Post {
	// 新代码：直接问数据库要数据
	return dao.GetAllPosts()
}

// CreatePost 发布文章 (为下一步做准备)
func (s *PostService) CreatePost(post *model.Post) {
	dao.CreatePost(post)
}

func (s *PostService) GetPostByID(id int) model.Post {
	return dao.GetPostByID(id)
}
