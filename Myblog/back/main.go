package main

import (
	"Myblog/dao"
	"Myblog/model"
	"Myblog/router"
	"fmt"
)

func main() {
	// 1. 初始化数据库
	dao.InitDB()

	// --- ⚡️ 临时代码：如果是第一次运行，插一条数据进去，不然页面是空的不好看 ---
	if len(dao.GetAllPosts()) == 0 {
		dao.CreatePost(&model.Post{
			Title:   "Hello Database!",
			Summary: "这是存储在 SQLite 数据库里的第一篇文章。",
			Content: "# 成功连接数据库\n\n恭喜你，现在你的博客数据已经持久化了！",
		})
		fmt.Println("已自动插入测试数据")
	}
	// -------------------------------------------------------------------

	// 2. 启动路由
	r := router.InitRouter()
	r.Run(":8080")
}
