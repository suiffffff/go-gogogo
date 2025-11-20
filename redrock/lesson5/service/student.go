package service

//Service则是用来调用dao的函数，起到一个转接的作用
import (
	"lesson5/dao"
	"lesson5/model"
	"lesson5/utils"
	"time"
)

func CreateStudent(name string, age int, grade string, rawPassword string) (*model.Student, error) {
	password := utils.Jiami(rawPassword)
	student := &model.Student{
		Name:     name,
		Age:      age,
		Grade:    grade,
		Password: password,
	}
	err := dao.CreateStudent(student)
	return student, err
}
func Login(name string, rawPassword string) (string, string, error) {
	jiami := utils.Jiami(rawPassword)
	student, err := dao.GetStudentLogin(name, jiami)
	if err != nil {
		return "", "", err
	}
	accessToken, refreshToken, err := utils.GenerateTokens(uint64(student.ID), student.Role)
	if err != nil {
		return "", "", err
	}
	exp := time.Now().Add(7 * 24 * time.Hour).Unix()
	if err := dao.StoreRefreshToken(uint(student.ID), refreshToken, exp); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}
func GetCourseList() ([]model.ClassInfo, error) {
	classes, err := dao.GetClassList()
	if err != nil {
		return nil, err
	}
	//测测人
	for i := range classes {
		classes[i].Left = classes[i].Capacity - classes[i].Current
	}
	return classes, nil
}
func GetMyCourseList(studentID uint64) ([]model.MyClassInfo, error) {
	classes, err := dao.GetMyCourse(studentID)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
func Xuanke(studentID int, xiaotuantiID int) error {
	return dao.EnrollStudent(int32(studentID), int32(xiaotuantiID))
}
func DropCourse(studentID int, classID int) error {
	return dao.DropCourse(studentID, classID)
}
