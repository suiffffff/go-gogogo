package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 网址，以后网址变更可以直接修改这里
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number int
	//具体问题网址
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login string
	//提问者个人github网址
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	//转义，因为只有基本字符能在URL上传输，为了表达字符需要先转义，否则无法识别
	q := url.QueryEscape(strings.Join(terms, " "))
	//模拟一个GET请求
	requestURL := IssuesURL + "?q=" + q
	fmt.Println("正在请求 URL:", requestURL)

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	//延迟关闭
	defer resp.Body.Close()
	//StatusOK就是200，代表访问成功
	if resp.StatusCode != http.StatusOK {
		//这里关闭的TCP网络连接
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	//resp。body是数据流，newdecode是将这些数据流解码成json的形式，decode是将数据开始传输进resul里面
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
func main() {
	// 这里搜索的是：golang/go 仓库中，包含 "json" 和 "decoder" 关键词的问题
	searchTerms := []string{"repo:golang/go", "json", "decoder"}
	fmt.Println("开始搜索 GitHub Issues...")
	result, err := SearchIssues(searchTerms)
	if err != nil {
		// 如果出错，记录日志并退出程序
		log.Fatalf("搜索出错: %v", err)
	}

	// 打印结果统计
	fmt.Printf("成功！总共找到 %d 个结果。\n", result.TotalCount)
	fmt.Println("------------------------------------------------------------------------")

	// 遍历并打印前 10 条（或者更少，如果结果不足10条）
	for i, item := range result.Items {
		if i >= 10 {
			break
		}
		// 格式化打印：
		// %-5d: 左对齐的数字，占5位 (Issue编号)
		// %-15.15s: 左对齐字符串，占15位，超过截断 (用户名)
		// %.50s: 字符串最长显示50个字符 (标题)
		fmt.Printf("#%-5d %-15.15s %.50s\n",
			item.Number, item.User.Login, item.Title)
	}
}

//可以去了解一下gin框架，之后对这个的理解应该更深
//这些代码是站在客户端的方向，gin框架是站在服务端的方向
//客户端是发送请求，gin是接受请求

//也可以做一个爬虫，我把它放在了Pro文件夹里面，学会的话应该会有更深的感触
