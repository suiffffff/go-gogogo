package api

import (
	"lesson5/dao"
	"lesson5/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCourseHandler(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(400, gin.H{
			"err": "86不存在的课程",
		})
		return
	}
	if err := service.AddCourse(name); err != nil {
		c.JSON(500, gin.H{
			"err": "我的幽灵课程",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":    "重邮里的宠物课程",
		"course": name,
	})
}
func DeleteCourseHandler(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	if err := dao.DeleteCourse(id); err != nil {
		c.JSON(500, gin.H{
			"err": "OVERLOAD不死者之课",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "可塑性课程",
	})
}
func UpdateTeacherHandler(c *gin.Context) {
	newName := c.PostForm("newName")
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	if err := dao.UpdataTeacher(id, newName); err != nil {
		c.JSON(500, gin.H{
			"err": "暗杀教师",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":     "未闻师名",
		"Newname": newName,
	})
}
