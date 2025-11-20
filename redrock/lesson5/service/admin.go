package service

import "lesson5/dao"

func AddCourse(name string) error {
	return dao.AddcCourse(name)
}
func DeleteCourse(id uint) error {
	return dao.DeleteCourse(id)
}
func UpdateTeacher(id uint, newName string) error {
	return dao.UpdataTeacher(id, newName)
}
func AddTeacher(name string) error {
	return dao.AddTeacher(name)
}
func CreateClass(courseID uint, teacherID uint, capacity int) error {
	return dao.CreateClass(courseID, teacherID, capacity)
}
func UpdateClassCapacity(classID uint, newCapacity int) error {
	return dao.UpdateClassCapacity(classID, newCapacity)
}
