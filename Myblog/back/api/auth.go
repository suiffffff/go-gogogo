package api

import (
	"Myblog/dao"
	"Myblog/model"
	"Myblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 1. 加密密码
	hash, _ := utils.HashPassword(password)

	// 2. 存入数据库
	user := model.User{Username: username, Password: hash}
	result := dao.DB.Create(&user)

	if result.Error != nil {
		c.JSON(200, gin.H{"msg": "注册失败，用户名可能已存在"})
		return
	}
	c.Redirect(http.StatusFound, "/login") // 注册成功跳去登录
}

// Login 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user model.User
	dao.DB.Where("username = ?", username).First(&user)

	// 1. 验证密码
	if user.ID == 0 || !utils.CheckPasswordHash(password, user.Password) {
		c.JSON(200, gin.H{"msg": "用户名或密码错误"})
		return
	}

	// 2. 设置 Cookie (最简单的登录维持方式)
	// 在浏览器存一个叫 "user_id" 的 cookie，有效期 3600秒
	c.SetCookie("user_id", strconv.Itoa(int(user.ID)), 3600, "/", "", false, true)

	c.Redirect(http.StatusFound, "/")
}

// Logout 登出
func Logout(c *gin.Context) {
	// 清除 cookie
	c.SetCookie("user_id", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/")
}
