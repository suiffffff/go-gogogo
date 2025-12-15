package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

// --- 配置区域 ---

// 申请GitHub的令牌以爬取信息
// 注意：千万不要把带有真实 Token 的代码上传到公开的 GitHub 仓库！
const GitHubToken = "#123456789"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number int
	Title  string
	User   *User
}

type User struct {
	Login string
}

// --- 核心爬虫函数 ---

// SearchIssues 现在多了一个 page 参数，用于指定抓第几页
func SearchIssues(terms []string, page int) (*IssuesSearchResult, error) {
	// 1. 构建 URL，增加 &page= 参数
	q := url.QueryEscape(strings.Join(terms, " "))
	// IssueURL是搜索地址，？可以看看gin框架的get，q代表搜索内容，per_page=100 表示一页抓100条（最大值），page=%d 是动态页码
	//sprintf是返回字符串，拼接字符，最主要的优点是能用占位符
	requestURL := fmt.Sprintf("%s?q=%s&per_page=100&page=%d", IssuesURL, q, page)

	// 2. 创建请求对象 (http.NewRequest)
	// Newrequest和get的不同，可以想象成fmt.Fprintf和fmt.printf
	//body是nil是因为是get请求，这里可以去了解一下post，这里就需要传json
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	// 3. 【关键步骤】添加身份认证 Header
	// 这就是让 GitHub 知道“是你”在访问
	if GitHubToken != "" {
		//这里Authorization是身份验证，也是key
		//后面的就是value
		//键值对，json的东西，主要是告诉服务器这个键对的什么值
		//authorization在GitHub这个网址代表身份验证，后面的令牌就是你给他验证的东西
		req.Header.Set("Authorization", "token "+GitHubToken)
	}
	// 这里是自我介绍，user-agent用户账户，告诉网站用户是谁，有的网址对爬虫的监管较严，需要告诉服务器用户
	req.Header.Set("User-Agent", "My-Go-Crawler/1.0")
	//人在国外可以用这个
	// 4. 发送请求
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	// 检查状态码
	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("请求失败: 状态码 %d", resp.StatusCode)
	//}
	// 4. 发送请求
	// 添加代理设置
	//因为国内容易连不上GitHub，所以需要走代理端
	// 设置你的代理地址 (看一下梯子的代理端)
	proxyUrl, _ := url.Parse("http://127.0.0.1:7897")

	client := &http.Client{
		// Transport 是 Go 网络库的底层机制，负责 TCP 连接
		Transport: &http.Transport{
			//// 告诉底层：所有流量都通过 proxyUrl（也就是你的梯子） 转发
			Proxy: http.ProxyURL(proxyUrl),
		},
		//设置超时
		Timeout: 10 * time.Second,
	}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//// 5. 解码
	//var result IssuesSearchResult
	//if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	//	return nil, err
	//}
	// 这里是调试代码
	// 1. 先把 Body 全部读出来变成字节
	bodyBytes, _ := io.ReadAll(resp.Body)

	// 2. 转成字符串打印
	bodyString := string(bodyBytes)
	fmt.Printf("【调试】状态码: %d\n", resp.StatusCode)
	fmt.Printf("【调试】Body内容: %s\n", bodyString)

	// 3. 如果 Body 是空的，或者不是 JSON，这里就能看出来
	if len(bodyString) == 0 {
		return nil, fmt.Errorf("GitHub 返回了空数据，可能是网络拦截")
	}

	// body是一个流，文件不能一下子全部打包在传输，是一点一点传输过来的，就像下载的进度条一样，被readall读取到里面的东西后，就不能够再次被读取了
	// 所以这里不能直接用 json.NewDecoder(resp.Body)，得把数据重新塞进去或者直接 Unmarshal
	// 为了简单，我们直接用 Unmarshal
	var result IssuesSearchResult
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("JSON 解析失败: %v, 内容: %s", err, bodyString)
	}
	return &result, nil
}

// --- 主程序 ---

func main() {
	searchTerms := []string{"repo:gin-gonic/gin", "router", "bug"}

	fmt.Println("启动爬虫...")
	fmt.Printf("使用的 Token: %s******\n", GitHubToken[:4]) // 为了安全只显示前4位

	page := 1
	totalItems := 0

	// 无限循环，直到没有数据为止
	for {
		fmt.Printf("正在抓取第 %d 页...\n", page)

		result, err := SearchIssues(searchTerms, page)
		if err != nil {
			log.Printf("抓取第 %d 页失败: %v\n", page, err)
			break // 出错就停止
		}

		// 如果这一页没有数据了 (Items 为空)，说明抓完了
		if len(result.Items) == 0 {
			fmt.Println("数据已全部抓取完毕！")
			break
		}

		// 处理这一页的数据
		for _, item := range result.Items {
			// 这里仅仅是打印，实际项目中你可能会把它们存入数据库
			fmt.Printf("[%d] %s (by %s)\n", item.Number, item.Title, item.User.Login)
			totalItems++
		}

		// 简单的防封策略：如果你没有 Token，必须加上这一行休息
		// 有 Token 也可以稍微休息一下，做个有礼貌的爬虫
		time.Sleep(1 * time.Second)

		// 准备抓下一页
		page++
	}

	fmt.Printf("----------------------------------\n")
	fmt.Printf("任务结束，总共抓取到了 %d 条数据。\n", totalItems)
}

//学会了这个，可以再进阶一点爬取个b站
