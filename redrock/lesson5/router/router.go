package router

//路由则是启动服务器并给出地址
import (
	"lesson5/api"
	"lesson5/middleware"

	"github.com/gin-gonic/gin"
)

// cmd
// curl -X POST "http://127.0.0.1:8080/api/create-student" -d "name=奶茶" -d "password=giveme" -d "age=99" -d "grade=haha"
// curl -X POST "http://127.0.0.1:8080/api/login" -d "name=权限狗" -d "password=123456"

// 登录的话感觉太麻烦了，目前没找到什么优化的方法，只能凑合用下面的
// powershell
// curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=权限狗" -d "password=123456"
// curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=选课老登" -d "password=654321"
// $token = (curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=权限狗" -d "password=123456" | ConvertFrom-Json).access_token
// $user = (curl.exe -X POST "http://127.0.0.1:8080/api/login" -d "name=选课老登" -d "password=654321" | ConvertFrom-Json).access_token
// curl.exe -X GET "http://127.0.0.1:8080/api/admin/check" -H "Authorization: Bearer $token"
// curl.exe -X GET "http://127.0.0.1:8080/api/admin/check" -H "Authorization: Bearer $user"

//curl.exe -X GET "http://127.0.0.1:8080/api/courses" -H "Authorization: Bearer $user"

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
		adminGroup.GET("/check", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "哟西，尊贵滴管理员，请进请进！"})
		})
	}
	return r
}
