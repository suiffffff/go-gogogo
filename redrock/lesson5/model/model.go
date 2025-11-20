package model

// model里面主要是存储数据结构，对于数据的加密不在这里
type Student struct {
	ID       int32
	Name     string
	Age      int
	Grade    string
	Password string
	Role     string
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
	StudentID    int32 `gorm:"primaryKey"`
	XiaotuantiID int32 `gorm:"primaryKey"`
}
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

type UserToken struct {
	ID uint `gorm:"primaryKey"`
	//索引，看介绍大概是类似于地址的一个东西吧？数据库应该是类似单链表？没有地址就需要遍历
	UserID uint `gorm:"index"`
	//强制类型转换后防止数据库插入两个一模一样的token
	Token string `gorm:"type:varchar(512);unique"`
	//这里存储过期时间
	ExpiresAt int64
	//这里是是否撤销token，默认false
	Revoked bool `gorm:"default:false"`
}

// 这里又到json环节了，不会，先用着吧
type ClassInfo struct {
	ClassID     int    `json:"class_id"`     // 小团体ID (选课用的ID)
	CourseName  string `json:"course_name"`  // 课程名
	TeacherName string `json:"teacher_name"` // 老师名
	Capacity    int    `json:"capacity"`     // 总容量
	Current     int    `json:"current"`      // 当前人数
	Left        int    `json:"left"`         // 剩余名额
}
type MyClassInfo struct {
	ClassID     int    `json:"class_id"`
	CourseName  string `json:"course_name"`
	TeacherName string `json:"teacher_name"`
}
