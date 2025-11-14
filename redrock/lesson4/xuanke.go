package main

// go get github.com/go-sql-driver/mysql
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql
//go get -u github.com/gin-gonic/gin

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    int32
	Name  string
	Age   int
	Grade string
}
type Course struct {
	ID   int32
	Name string
}
type Teacher struct {
	ID   int32
	Name string
}
type Redrockclass struct {
	//这里是关联表？所有关联建必须用后面那个gorm。但说实话我也不知道是什么意思。只能先记住
	StudentID    int32 `gorm:"primaryKey"`
	XiaotuantiID int32 `gorm:"primaryKey"`
}

// Xiaotuanti
// 我这里确实搞不懂，ai说的也含糊，ai最后含糊的说gorm会选择性修复不匹配的项
// 我确实不太明白。为什么前面的项可以不管，这里的表就必须修复
// 根据我的进一步测试，两个int类型不匹配会被无视，三个以上会报错，所以到底是什么逻辑？搞不懂
type Xiaotuanti struct {
	ID        int32 `gorm:"primaryKey"`
	CourseID  int32
	TeacherID int32
	Capacity  int32
	Current   int32
}

func (Xiaotuanti) TableName() string {
	return "xiaotuantis"
}

// 我说为什么会错，表变成xiao_tuanti了草，这里修改返回名

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(mysql.Open(
		"root:Paow7778200400@@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local",
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("控制失败")
		panic(err)
	}
	fmt.Println("我真得控制你咯")
	err = db.AutoMigrate(&Student{}, &Course{}, &Teacher{}, &Redrockclass{}, &Xiaotuanti{})
	if err != nil {
		//这里是强行退出，和panic不一样，panic会执行defer，但是fatal不会
		log.Fatal("自动迁移失败:", err)
	}
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/xuanke", xuanke)
		api.GET("/create-student", create)
	}
	//这里我真不到啊，用static一直在报路径的错，我也看不懂错哪里了，只好先用ai换了个函数
	//将网页由静态转动态需要JavaScript，这个我是真一点都不知道，就先不写了
	r.Use(static.Serve("/", static.LocalFile("./index", true)))
	fmt.Println("选课，启动！")
	r.Run()
}
func create(c *gin.Context) {
	name := c.Query("name")
	agestr := c.Query("age")
	grade := c.Query("grade")
	if name == "" || agestr == "" || grade == "" {
		c.JSON(400, gin.H{
			"err": "不像个人",
		})
		return
	}
	age, err := strconv.Atoi(agestr)
	if err != nil {
		c.JSON(400, gin.H{
			"err": "不像个数字",
		})
		return
	}
	student := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	if err := db.Create(&student).Error; err != nil {
		c.JSON(400, gin.H{
			"err": "反正就是不行，爱咋咋吧",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "可以选学姐咯！",
		"data":    student, // 把新创建的学生信息返回
	})
}
func xuanke(c *gin.Context) {
	studentIDStr := c.Query("studentid")
	xiaotuantiIDStr := c.Query("xiaotuantiid")
	_studentID, err1 := strconv.Atoi(studentIDStr)
	_xiaotuantiID, err2 := strconv.Atoi(xiaotuantiIDStr)
	if err1 != nil {
		c.JSON(400, gin.H{
			"err": "你是啥学生",
		})
		return
	}
	if err2 != nil {
		c.JSON(400, gin.H{
			"err": "你想选啥课",
		})
		return
	}
	studentID := int32(_studentID)
	xiaotuantiID := int32(_xiaotuantiID)
	err := db.Transaction(func(tx *gorm.DB) error {
		//这里解决了一个并发的问题，ai说if的话同时访问会有问题，为了严谨还是加上了
		//只能认得找找课了
		var xiaotuanti Xiaotuanti
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&xiaotuanti, xiaotuantiID).Error; err != nil {
			//返回err确实是可以，但是不会知道错在哪，于是问了问ai，errors.New能顺便回个报错信息，或许还有更好的方法？
			return errors.New("锁失败")
		}
		if xiaotuanti.Current >= xiaotuanti.Capacity {
			return errors.New("课程已满")
		}
		joinRecord := Redrockclass{StudentID: studentID, XiaotuantiID: xiaotuantiID}
		if err := tx.Create(&joinRecord).Error; err != nil {
			return errors.New("选课失败")
		}
		if err := tx.Exec("update xiaotuantis set current = current+1 where id = ?", xiaotuantiID).Error; err != nil {
			return errors.New("更新失败")
		}
		return nil
	})
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

//之前一直1452，找了半天问题，最后一不小心删了redrockclasses，发现xiaotuanti不存在，因为我之前用过，为了对应把xiaotuanti删了加了个xiaotuantis
//然后redrockclasses存了个错误的信息，ai半天没发现，我也没发现
//自己确实比较粗心的吧，这个错浪费一个多小时
//以及，这部分代码实际上只能做到理解，真要说重新写一次的话，估计是写不出来的
//但大部分信息仍然是靠ai获取，改错改了半天最后还加了一堆int32去解决1452的问题
//同样的，代码如果有问题的地方望指正，因为我确实不清楚ai到底在哪方面会给我错误信息
//我只能调试确定代码确实能跑，以及理解意思
//如上
