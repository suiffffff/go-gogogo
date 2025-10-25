package main

import (
	"fmt"
	"os"
)

func main() {
	//s, sep := "", ""
	//range可以遍历数组或者切片，然后像前面传参（应该是传值吧？）idx接受的行数（索引），arg接受的内容（参考值），[a:b]是指从数组a遍历到b（不写表示从头/尾）（左闭右开）
	//需要注意的是，range自动包含终止条件，即遍历完终止
	for idx, arg := range os.Args[1:] {
		/*s += sep + arg
		sep = " "*/
		fmt.Printf("切片索引：%d，参数值：%s\n", idx, arg)
	}
	//fmt.Println(s)
}
