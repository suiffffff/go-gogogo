package api

//api则是解析http的参数，然后把读取到的数据传入给service
import (
	"lesson5/service"
	"lesson5/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateStudentHandler(c *gin.Context) {
	//众所周知get的话密码会被查浏览记录，肯定不太安全
	//post是隐蔽的存储数据，postform才是拿数据
	//json的话，不知道捏
	name := c.PostForm("name")
	agestr := c.PostForm("age")
	grade := c.PostForm("grade")
	password := c.PostForm("password")
	if name == "" || agestr == "" || grade == "" {
		c.JSON(400, gin.H{"err": "不像个人(参数不全)"})
		return
	}
	if password == "" {
		c.JSON(400, gin.H{"err": "你（密）吗呢？"})
		return
	}
	age, err := strconv.Atoi(agestr)
	if err != nil {
		c.JSON(400, gin.H{"err": "不像个数字"})
		return
	}
	student, err := service.CreateStudent(name, age, grade, password)
	if err != nil {
		c.JSON(400, gin.H{"err": "反正就是不行，爱咋咋吧"})
		return
	}
	c.JSON(200, gin.H{
		"message": "可以选学姐咯！",
		"data":    student,
	})
}
func LoginHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		c.JSON(400, gin.H{
			"err": "名字和密码，总有一个神隐",
		})
		return
	}
	at, rt, err := service.Login(name, password)
	//这里说是数据库异常。token生成失败，数据异常。AI说模糊信息，但这些黑客能修改吗？
	if err != nil {
		c.JSON(400, gin.H{
			"err": "你码错咯",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":           "登陆成功",
		"access_token":  at,
		"refresh_token": rt,
	})
}
func RefreshTokenHandler(c *gin.Context) {
	refreshTokenStr := c.PostForm("refresh_token")
	if refreshTokenStr == "" {
		c.JSON(400, gin.H{"err": "你码呢"})
		return
	}
	claims, err := utils.VerifyRefreshToken(refreshTokenStr)
	if err != nil {
		c.JSON(401, gin.H{"err": "你码不对或过期咯"})
		return
	}
	isValid, _ := service.CheckRefreshToken(refreshTokenStr)
	if !isValid {
		c.JSON(401, gin.H{"err": "你被ban咯"})
		return
	}
	newAccess, newRefresh, err := utils.GenerateTokens(claims.UserID, claims.Role)
	if err != nil {
		c.JSON(500, gin.H{"err": "系统错误，生成你码失败"})
		return
	}
	c.JSON(200, gin.H{
		"msg":           "你新码来咯",
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}
func GetCourseListHandler(c *gin.Context) {
	list, err := service.GetCourseList()
	if err != nil {
		c.JSON(500, gin.H{
			"err": "世界上最悲哀的事，莫过于不想选课，不得不选，又选不上",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "天王降临",
		"data": list,
	})
}
func XuankeHandler(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.JSON(400, gin.H{"err": "未登录"})
		return
	}
	StudentID := id.(uint)
	xiaotuantiIDStr := c.PostForm("xiaotuantiid")
	xiaotuantiID, err := strconv.Atoi(xiaotuantiIDStr)
	if err != nil {
		c.JSON(400, gin.H{"err": "你想选啥课(ID错误)"})
		return
	}
	XiaotuantiID := uint(xiaotuantiID)
	err = service.Xuanke(StudentID, XiaotuantiID)
	if err != nil {
		c.JSON(400, gin.H{
			"(选课）启动": "失败",
			"err":    err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"（选课）启动":  "成功",
			"message": "你已成功选课，快来骚扰学姐吧",
		})
	}
}
func GetMyCourseListHandler(c *gin.Context) {
	id, exists := c.Get("userID")
	if !exists {
		c.JSON(500, gin.H{
			"err": "用户在哪里",
		})
		return
	}
	//这里不知道为什么uint（）不行
	ID := id.(uint)
	list, err := service.GetMyCourseList(ID)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "你觉得你报上了吗",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "你知道的，李海军一直是cy最好的老师，我推荐去报它的课程，我和它的化学反应好的不可思议",
		"data": list,
	})
}
func DropCourseHandler(c *gin.Context) {
	id, _ := c.Get("userID")
	ID := id.(uint)
	classIDStr := c.PostForm("class_id")
	if classIDStr == "" {
		c.JSON(400, gin.H{"err": "请告诉我要退哪门课"})
		return
	}
	classID, _ := strconv.Atoi(classIDStr)
	ClassID := uint(classID)
	if err := service.DropCourse(ID, ClassID); err != nil {
		c.JSON(400, gin.H{"err": "退课失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "我来到不是时候？"})
}
