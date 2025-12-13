package api

import (
	"Myblog/model"
	"Myblog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostAPI struct {
	PostService service.PostService
}

// CreatePost 处理表单提交
func (api *PostAPI) CreatePost(c *gin.Context) {
	// 1. 获取表单数据
	title := c.PostForm("title")
	content := c.PostForm("content")
	summary := c.PostForm("summary")

	// 2. 组装数据模型
	post := model.Post{
		Title:   title,
		Content: content,
		Summary: summary,
	}

	// 3. 调用 Service 保存到数据库
	api.PostService.CreatePost(&post)

	// 4. 保存成功后，跳转回首页看结果
	c.Redirect(http.StatusFound, "/")
}
