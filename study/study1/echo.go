package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//os.Args 是终端输入的，go run 文件名.go 之后的值，os.Args是切片，感觉切片其实就是c语言中的指针+动态内存管理的数组？就先当数组吧
	//os.Args 中的第0个元素就是文件路径，所以一般会略过
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//s能输出拼接后的值
	fmt.Println(s)
}
