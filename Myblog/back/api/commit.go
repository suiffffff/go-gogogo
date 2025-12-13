package api

import (
	"Myblog/dao"
	"Myblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddComment 发表评论
func AddComment(c *gin.Context) {
	content := c.PostForm("content")
	postID, _ := strconv.Atoi(c.PostForm("post_id"))

	// 从中间件获取当前登录用户 ID
	userIDStr, _ := c.Cookie("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	comment := model.Comment{
		Content: content,
		PostID:  uint(postID),
		UserID:  uint(userID),
	}
	dao.DB.Create(&comment)

	c.Redirect(http.StatusFound, "/post/"+c.PostForm("post_id"))
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	// 这里应该加一个判断：只有评论作者或管理员才能删除 (此处省略简化)
	dao.DB.Delete(&model.Comment{}, id)

	// 简单的返回上一页
	c.Redirect(http.StatusFound, "/")
}
