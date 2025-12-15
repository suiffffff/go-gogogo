package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 换汤不换药，对吧？
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//所以这东西为什么会被删了啊
		//ioutil显然是是一个库，readfile显然是读取文件内容，c中fopen后面跟个r？应该是
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		//这里的是string库里面的切割，作用是分割字符串，那么如何分割呢？按照后面的字符串分割
		//只要有后面的字符串的内容，就会完成一次分割，例如“a”，有a就会被分割，分割的不保留
		//需要注意的是，这里应该是动态内存做的字符串数组？不得而知
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s", n, line)
		}
	}

}
