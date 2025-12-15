package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

// --- 1. 数据结构定义 ---
// B站的 JSON 结构比 GitHub 稍微复杂一点点，是多层包裹的

type BiliResponse struct {
	Code    int       `json:"code"`    // 0 代表成功，其他代表失败
	Message string    `json:"message"` // 错误信息
	Data    *BiliData `json:"data"`
}

type BiliData struct {
	Result []*BiliVideo `json:"result"` // 搜索结果列表
}

type BiliVideo struct {
	Title    string `json:"title"`    // 视频标题（里面可能包含 HTML 标签）
	Author   string `json:"author"`   // UP主名字
	Play     int    `json:"play"`     // 播放量
	Bvid     string `json:"bvid"`     // 视频唯一ID (BV号)
	Duration string `json:"duration"` // 时长
}

// --- 2. 爬虫核心函数 ---

func SearchBilibili(keyword string, page int) (*BiliResponse, error) {
	// B站搜索接口 API
	const SearchAPI = "https://api.bilibili.com/x/web-interface/search/type"

	// 构建参数
	params := url.Values{}
	params.Set("search_type", "video") // 指定只搜视频
	params.Set("keyword", keyword)
	params.Set("page", fmt.Sprintf("%d", page))
	params.Set("page_size", "50")
	// 拼接完整 URL
	requestURL := fmt.Sprintf("%s?%s", SearchAPI, params.Encode())
	fmt.Println("正在请求:", requestURL)

	// 创建请求
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	// 【关键伪装】 B站查得比较严的两个头
	// 1. 假装我是 Chrome 浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	// 2. 假装我是从B站官网点过来的 (防盗链检查)
	req.Header.Set("Referer", "https://www.bilibili.com/")
	//这里需要找找自己的Cookie，通过网站F12找到标头里的ookie然后复制粘贴
	req.Header.Set("Cookie", "#")
	// 发送请求
	// 因为B站通常不需要代理（除非你在国外），所以这里用默认 Client 即可
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 状态码错误: %d", resp.StatusCode)
	}

	// 解码
	var result BiliResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// B站特有的逻辑：HTTP 200 不代表成功，要看 body 里的 code 字段
	if result.Code != 0 {
		return nil, fmt.Errorf("B站业务报错: code=%d, message=%s", result.Code, result.Message)
	}

	return &result, nil
}

// --- 3. 主程序 ---

func main() {
	keyword := "cos"
	maxPage := 100
	var allVideos []*BiliVideo
	fmt.Printf("开始搜索: %s，计划抓取 %d 页...\n", keyword, maxPage)
	fmt.Println("---------------------------------------------------")
	for page := 1; page <= maxPage; page++ {
		fmt.Printf("正在抓取第 %d 页... ", page)

		// 1. 发起请求
		resp, err := SearchBilibili(keyword, page)
		if err != nil {
			log.Printf("失败: %v\n", err)
			break // 出错就停止
		}

		// 2. 检查是否有数据
		// 如果 Data 是空的，或者 Result 长度为 0，说明B站没视频给咱们了
		if resp.Data == nil || len(resp.Data.Result) == 0 {
			fmt.Println("没有更多数据了，停止抓取。")
			break
		}

		// 3. 把这一页的 50 个视频，倒进我们的大仓库 (allVideos) 里
		// append(a, b...) 的意思是把 b 列表里的元素一个个追加到 a 后面
		allVideos = append(allVideos, resp.Data.Result...)

		fmt.Printf("成功抓到 %d 条数据。\n", len(resp.Data.Result))

		// 4. 【关键】休息一下！
		// 这一步非常重要，防止请求太快触发 412 封锁
		time.Sleep(2 * time.Second)
	}

	fmt.Println("---------------------------------------------------")
	fmt.Printf("任务结束！总共收集到 %d 个视频。\n", len(allVideos))

	fmt.Println("正在按播放量从高到低排序...")

	// sort.Slice 接收两个参数：
	// 参数1: 要排序的切片 (allVideos)
	// 参数2: 一个匿名函数，告诉 Go 怎么比较第 i 个和第 j 个元素
	sort.Slice(allVideos, func(i, j int) bool {
		// 如果你想要【从高到低】(降序)，就用 大于号 >
		// 意思是：如果前一个(i) 比 后一个(j) 大，那么前一个排在前面
		return allVideos[i].Play > allVideos[j].Play

		// 如果你想要【从低到高】(升序)，就改用 小于号 <
	})

	// --- 【3. 打印排序后的前 20 个】 ---
	fmt.Println("---------------------------------------------------")
	fmt.Println("【播放量排行榜 Top 20】")

	for i, video := range allVideos {
		if i >= 20 {
			break
		}

		cleanTitle := strings.ReplaceAll(video.Title, "<em class=\"keyword\">", "")
		cleanTitle = strings.ReplaceAll(cleanTitle, "</em>", "")

		// 打印时加上播放量，看看是不是真的排好了
		fmt.Printf("[%s] No.%-2d 播放:%-8d UP:%-10s %s\n", video.Bvid, i+1, video.Play, video.Author, cleanTitle)
	}
}
