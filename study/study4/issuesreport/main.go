package main

import (
	"log"
	"os"
	"study4/github"
)

func main() {
	// 1. 获取命令行参数（比如你在终端输入了 repo:golang/go json）
	// 如果没有参数，默认给一组，防止报错
	terms := os.Args[1:]
	if len(terms) == 0 {
		terms = []string{"repo:golang/go", "is:open", "json", "decoder"}
	}

	// 2. 调用你写好的 SearchIssues 函数获取数据
	// 注意：如果 SearchIssues 和 main 在同一个包，直接调用。
	// 如果在不同包，需要 import 并用 github.SearchIssues 调用。
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}

	// 使用 issueList.Execute 将数据填充进 HTML 模板
	// os.Stdout 表示直接输出到屏幕（稍后我们在命令行里重定向到文件）
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

//go run . "repo:golang/go" "json" "decoder" > issues.注意这是一个非标准库，需要下载
