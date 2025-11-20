package main

//main只需要启动各个板块
import (
	"fmt"
	"lesson5/dao"
	"lesson5/router"
)

func main() {
	dao.InitMySQL()
	fmt.Println("我真得控制你咯 (数据库连接成功)")

	r := router.SetupRouter()

	fmt.Println("选课，启动！")
	r.Run()
}
