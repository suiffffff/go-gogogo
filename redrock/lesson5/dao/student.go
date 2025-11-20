package dao

//这一部分为对数据库的增删查改，涉及到的函数应该都写入这里
import (
	"errors"
	"lesson5/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateStudent(student *model.Student) error {
	return DB.Create(student).Error
}
func EnrollStudent(studentID int32, xiaotuantiID int32) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var xiaotuanti model.Xiaotuanti
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&xiaotuanti, xiaotuantiID).Error; err != nil {
			return errors.New("锁失败或课程不存在")
		}
		if xiaotuanti.Current >= xiaotuanti.Capacity {
			return errors.New("课程已满")
		}
		joinRecord := model.Redrockclass{StudentID: studentID, XiaotuantiID: xiaotuantiID}
		if err := tx.Create(&joinRecord).Error; err != nil {
			return errors.New("选课失败，可能重复选课")
		}
		if err := tx.Exec("update xiaotuantis set current = current+1 where id = ?", xiaotuantiID).Error; err != nil {
			return errors.New("更新人数失败")
		}
		return nil
	})
}
func GetStudentLogin(name string, password string) (*model.Student, error) {
	var student model.Student
	err := DB.Where("name=? AND password=?", name, password).First(&student).Error
	return &student, err
}
func GetClassList() ([]model.ClassInfo, error) {
	var classes []model.ClassInfo
	err := DB.Table("xiaotuantis").
		//拼好表
		Select("xiaotuantis.id as class_id, courses.name as course_name, teachers.name as teacher_name, xiaotuantis.capacity, xiaotuantis.current").
		//你去把courses的id偷过来
		Joins("left join courses on xiaotuantis.course_id = courses.id").
		//你去把teachers的id偷过来
		Joins("left join teachers on xiaotuantis.teacher_id = teachers.id").
		//插入classes
		Scan(&classes).Error
	return classes, err
}
func GetMyCourse(studentID uint64) ([]model.MyClassInfo, error) {
	var myClasses []model.MyClassInfo
	//主表应该是要找到关键的对应消息，就像之前查询课程，对于课程什么最关键呢？课的ID，老师，还有课程名，这些在哪里被关联了呢？是小团体这个表，我们要顺着这个表往下去找，而不是从下往上去找
	//所以主表都应该是一个关联表，顺关联表往下去找其他表
	err := DB.Table("redrockclasses").
		Select("xiaotuantis.id as class_id, courses.name as course_name, teachers.name as teacher_name").
		Joins("left join xiaotuantis on redrockclasses.xiaotuanti_id = xiaotuantis.id").
		Joins("left join courses on xiaotuantis.course_id = courses.id").
		Joins("left join teachers on xiaotuantis.teacher_id = teachers.id").
		Where("redrockclasses.student_id = ?", studentID).
		Scan(&myClasses).Error
	return myClasses, err
}
func DropCourse(studentID int, classID int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("student_id = ? AND xiaotuanti_id = ?", studentID, classID).Delete(&model.Redrockclass{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("课都没选捏")
		}
		if err := tx.Model(&model.Xiaotuanti{}).Where("id = ?", classID).UpdateColumn("current", gorm.Expr("current - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}
