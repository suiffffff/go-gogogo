package api

import (
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
	ID := uint(id)
	if err := service.DeleteCourse(ID); err != nil {
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
	ID := uint(id)
	if err := service.UpdateTeacher(ID, newName); err != nil {
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
func AddTeacherHandler(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(400, gin.H{
			"err": "青春猪头少年做梦不会梦到兔女郎学姐",
		})
		return
	}

	if err := service.AddTeacher(name); err != nil {
		c.JSON(500, gin.H{
			"err": "添加失败: " + err.Error(),
		})
		return
	}
	c.JSON(200,
		gin.H{"msg": "师者无敌",
			"teacher": name,
		})
}

func CreateClassHandler(c *gin.Context) {
	courseIDStr := c.PostForm("course_id")
	teacherIDStr := c.PostForm("teacher_id")
	capacityStr := c.PostForm("capacity")

	// 类型转换
	cID, _ := strconv.Atoi(courseIDStr)
	tID, _ := strconv.Atoi(teacherIDStr)
	capacity, _ := strconv.Atoi(capacityStr)
	err := service.CreateClass(uint(cID), uint(tID), capacity)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "开课失败: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "小葵花课堂开课咯！学生学习老不认真...",
	})
}

func UpdateCapacityHandler(c *gin.Context) {
	classIDStr := c.PostForm("class_id")
	newCapStr := c.PostForm("capacity")
	classID, _ := strconv.Atoi(classIDStr)
	newCap, _ := strconv.Atoi(newCapStr)
	if err := service.UpdateClassCapacity(uint(classID), newCap); err != nil {
		c.JSON(500, gin.H{
			"err": "就这吗？",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "嗯~好大~",
	})
}
