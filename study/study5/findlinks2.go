package main

import (
	"fmt"
	"net/http"
	"os"
	"study5/visit"

	"golang.org/x/net/html"
)

func main() {
	//很熟悉的一个函数嘛，获取命令行之后的东西
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// 接受传来的字符串
func findLinks(url string) ([]string, error) {
	//get是尝试访问网址，会返回常见的网站结构体，我们后面细说
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//这里是状态码，ok表示成功访问，状态码常见的就是404（not found），这样是不是明白了一些？这里用200替换ok也是可以的
	if resp.StatusCode != http.StatusOK {
		//关闭网站
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	//这里是将网站的内容传给parse
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit.Visit(nil, doc), nil
}

//go run findlinks2.go https://golang.google.cn
