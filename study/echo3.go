package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//string.Join是一个拼接函数（含两个参数，第一个和第二个按顺序插入，类似于c语言的strcat？
	fmt.Println(strings.Join(os.Args[1:], " "))
}
