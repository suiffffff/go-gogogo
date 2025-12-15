package main

import (
	"fmt"
	"os"
	"study5/visit"

	"golang.org/x/net/html"
)

func main() {
	//os.stdin显然是标准输入，将输入的放进parse这个函数里处理
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit.Visit(nil, doc) {
		fmt.Println(link)
	}
}

//终端输入
//curl.exe https://golang.google.cn | go run findlinks1.go
