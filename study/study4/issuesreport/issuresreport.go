package main

//这里主要是展示作用
import (
	"log"
	"os"
	"text/template"
	"time"
)

// 1. 定义与模板对应的数据结构
// 这里的字段名首字母必须大写（导出），否则模板无法读取
type User struct {
	Login string
}

type Item struct {
	Number    int
	User      User
	Title     string
	CreatedAt time.Time
}

type IssueReport struct {
	TotalCount int
	Items      []Item
}

// 2. 自定义函数：计算时间是几天前
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// 书中的模板字符串
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

// 为了调用函数将main改成了oldmain
func oldmain() {
	// 3. 准备模拟数据
	report := IssueReport{
		TotalCount: 2,
		Items: []Item{
			{
				Number:    1024,
				User:      User{Login: "xiao_ming"},
				Title:     "Fix the critical bug in production database",
				CreatedAt: time.Now().Add(-24 * 3 * time.Hour), // 3天前
			},
			{
				Number: 1025,
				User:   User{Login: "li_hua"},
				// 这个标题故意弄得很长，测试模板里的 %.64s 截断功能
				Title:     "这是一个非常非常长的标题，用来测试模板里的 printf 管道命令是否能正常工作，如果工作正常，这段文字会被截断...",
				CreatedAt: time.Now().Add(-24 * 10 * time.Hour), // 10天前
			},
		},
	}

	// 4. 创建并解析模板
	// 关键点：必须在 Parse 之前使用 Funcs 注册自定义函数
	// "daysAgo" 是我们在模板里写的名字，daysAgo 是 Go 语言里的函数名

	///template是模板的意思，所以这里是在干什么呢？new一个template嘛，当然是创建模板了，后面就是名字
	//模板是与主程序隔离的，也就是根本看不到你写的函数的啦，map是映射，funcmap很明显嘛，map[string]func，这样子是不是很好理解啦？
	//parse就是解析的意思，有了模板，有了暗语（也就是funcmap）下面是不是就要开始抄答案了？
	t, err := template.New("issueList").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)

	if err != nil {
		log.Fatal(err)
	}

	// 5. 执行模板，将 report 数据填充进去，输出到标准输出 (os.Stdout)

	//err := t.Execute(输出目的地, 数据源)
	//显而易见的输出程序嘛
	if err := t.Execute(os.Stdout, report); err != nil {
		log.Fatal(err)
	}
}
