package dao

import (
	"lesson5/model"
)

func AddcCourse(name string) error {
	course := model.Course{Name: name}
	return DB.Create(&course).Error

}
func DeleteCourse(id uint) error {
	return DB.Delete(&model.Course{}, id).Error
}
func UpdataTeacher(id uint, newName string) error {
	return DB.Model(&model.Teacher{}).Where("id=?", id).Update("name", newName).Error
}
func AddTeacher(name string) error {
	teacher := model.Teacher{Name: name}
	return DB.Create(&teacher).Error
}
func CreateClass(courseID uint, teacherID uint, capacity int) error {
	xiaotuanti := model.Xiaotuanti{
		CourseID:  courseID,
		TeacherID: teacherID,
		Capacity:  capacity,
		Current:   0,
	}
	return DB.Create(&xiaotuanti).Error
}
func UpdateClassCapacity(classID uint, newCapacity int) error {
	return DB.Model(&model.Xiaotuanti{}).Where("id = ?", classID).Update("capacity", newCapacity).Error
}
