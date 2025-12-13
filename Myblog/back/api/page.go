package api

import (
	"Myblog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv" // 记得导入这个，用于把字符串转数字
)

type PageAPI struct {
	PostService service.PostService
}

// Index 渲染首页 (保持不变)
func (api *PageAPI) Index(c *gin.Context) {
	posts := api.PostService.GetPostList()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "首页",
		"posts": posts,
	})
}

// 🆕 新增：GetPostDetail 渲染详情页
func (api *PageAPI) GetPostDetail(c *gin.Context) {
	// 1. 获取 URL 中的 id 参数 (例如 /post/1 中的 "1")
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr) // 转成数字

	// 2. 调用 Service 查库
	post := api.PostService.GetPostByID(id)

	// 3. 渲染页面
	c.HTML(http.StatusOK, "post.html", gin.H{
		"title": post.Title,
		"post":  post,
	})
}
