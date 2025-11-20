package router

//路由则是启动服务器并给出地址
import (
	"lesson5/api"
	"lesson5/middleware"

	"github.com/gin-gonic/gin"
)

// cmd
// curl -X POST "http://127.0.0.1:8080/api/create-student" -d "name=奶茶" -d "password=giveme" -d "age=99" -d "grade=haha"
// curl -X POST "http://127.0.0.1:8080/api/create-student" -d "name=权限狗" -d "password=123456" -d "age=18" -d "grade=大一"

// curl -X POST "http://127.0.0.1:8080/api/login" -d "name=权限狗" -d "password=123456"

// 登录的话感觉太麻烦了，目前没找到什么优化的方法，只能凑合用下面的
// powershell
// curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=权限狗" -d "password=123456"
// curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=选课老登" -d "password=654321"

// $token = (curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=权限狗" -d "password=123456" | ConvertFrom-Json).access_token
// $user = (curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=选课老登" -d "password=654321" | ConvertFrom-Json).access_token

//[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
//$refresh = "你的refresh_token字符串"
//curl.exe -X POST "http://127.0.0.1:8080/api/refresh" -d "refresh_token=$refresh"
//$token = (curl.exe -X POST "http://127.0.0.1:8080/api/refresh" -d "refresh_token=$refresh" | ConvertFrom-Json).access_token
//$user = (curl.exe -X POST "http://127.0.0.1:8080/api/refresh" -d "refresh_token=$refresh" | ConvertFrom-Json).access_token

//验证管理员和学生

// curl.exe -X GET "http://127.0.0.1:8080/api/admin/check" -H "Authorization: Bearer $token"
// curl.exe -X GET "http://127.0.0.1:8080/api/admin/check" -H "Authorization: Bearer $user"

//管理员权限

//curl.exe -X POST "http://127.0.0.1:8080/api/admin/add-course" -H "Authorization: Bearer $token" -d "name=红颜怪谈"
//curl.exe -X POST "http://127.0.0.1:8080/api/admin/add-teacher" -H "Authorization: Bearer $token" -d "name=小男娘"
//curl.exe -X POST "http://127.0.0.1:8080/api/admin/create-class" -H "Authorization: Bearer $token" -d "course_id=6" -d "teacher_id=6" -d "capacity=50"
//curl.exe -X POST "http://127.0.0.1:8080/api/admin/update-capacity" -H "Authorization: Bearer $token" -d "class_id=6" -d "capacity=200"
//curl.exe -X GET "http://127.0.0.1:8080/api/courses" -H "Authorization: Bearer $user"

//学生权限

// curl.exe -X GET "http://127.0.0.1:8080/api/courses" -H "Authorization: Bearer $user"
// curl.exe -X POST "http://127.0.0.1:8080/api/xuanke" -H "Authorization: Bearer $user" -d "xiaotuantiid=6"
// curl.exe -X GET "http://127.0.0.1:8080/api/my-courses" -H "Authorization: Bearer $user"
// curl.exe -X POST "http://127.0.0.1:8080/api/drop-course" -H "Authorization: Bearer $user" -d "class_id=6"
func SetupRouter() *gin.Engine {
	r := gin.Default()
	public := r.Group("/api")
	{
		public.POST("/create-student", api.CreateStudentHandler)
		public.POST("/login", api.LoginHandler)
		public.POST("/refresh", api.RefreshTokenHandler)
	}

	userGroup := r.Group("/api", middleware.JWTAuthMiddleware())
	{
		userGroup.POST("/xuanke", api.XuankeHandler)
		userGroup.GET("/courses", api.GetCourseListHandler)
		userGroup.GET("/my-courses", api.GetMyCourseListHandler)
		userGroup.POST("/drop-course", api.DropCourseHandler)
	}
	adminGroup := r.Group("/api/admin", middleware.JWTAuthMiddleware(), middleware.AdminAuthMiddleware())
	{
		adminGroup.POST("/add-course", api.AddCourseHandler)
		adminGroup.POST("/delete-course", api.DeleteCourseHandler)
		adminGroup.POST("/update-teacher", api.UpdateTeacherHandler)
		adminGroup.POST("/add-teacher", api.AddTeacherHandler)
		adminGroup.POST("/create-class", api.CreateClassHandler)
		adminGroup.POST("/update-capacity", api.UpdateCapacityHandler)
		adminGroup.GET("/check", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "哟西，尊贵滴管理员，请进请进！"})
		})
	}
	return r
}
