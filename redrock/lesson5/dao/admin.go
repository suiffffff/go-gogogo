package dao

import (
	"lesson5/model"
)

func AddcCourse(name string) error {
	course := model.Course{Name: name}
	return DB.Create(&course).Error

}
func DeleteCourse(id int) error {
	return DB.Delete(model.Course{}, id).Error
}
func UpdataTeacher(id int, newName string) error {
	return DB.Model(model.Teacher{}).Where("id=?", id).Update("name=?", newName).Error
}
