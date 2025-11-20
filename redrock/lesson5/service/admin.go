package service

import "lesson5/dao"

func AddCourse(name string) error {
	return dao.AddcCourse(name)
}
func DeleteCourse(id int) error {
	return dao.DeleteCourse(id)
}
func UpdateTeacher(id int, newName string) error {
	return dao.UpdataTeacher(id, newName)
}
