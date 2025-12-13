package router

import (
	"Myblog/api"
	"Myblog/middleware" // 👈 别忘了引入你写的中间件包
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 1. 加载资源
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	// 2. 初始化需要实例化的 API (对应 api/page.go 和 api/post.go)
	pageAPI := api.PageAPI{}
	postAPI := api.PostAPI{}

	// --- 🌍 全局中间件 ---
	// 这一步非常重要！它会尝试从 Cookie 里读取 user_id。
	// 这样无论你在哪个页面，前端都能知道“当前是谁登录了”，从而显示“登录/注销”按钮。
	r.Use(middleware.AuthMiddleware())

	// --- 🔓 公开页面 (View) ---
	r.GET("/", pageAPI.Index)
	r.GET("/post/:id", pageAPI.GetPostDetail)

	// 登录 & 注册页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{"title": "登录"})
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{"title": "注册"})
	})

	// --- 🚀 公开接口 (API) ---
	// 对应 api/auth.go 里的函数
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.GET("/logout", api.Logout)

	// --- 🔒 隐私区域 (需要登录才能操作) ---
	// 使用 ForceLogin 中间件：如果没有登录，直接跳回 /login 页面
	authorized := r.Group("/", middleware.ForceLogin())
	{
		// 1. 评论相关 (api/comment.go)
		authorized.POST("/comment", api.AddComment)
		authorized.GET("/comment/delete/:id", api.DeleteComment)

		// 2. 写文章相关 (原来的 BasicAuth 被移除，改为用统一的用户登录保护)
		authorized.GET("/publish", func(c *gin.Context) {
			c.HTML(200, "publish.html", gin.H{"title": "发布文章"})
		})
		authorized.POST("/post", postAPI.CreatePost)
	}

	return r
}
