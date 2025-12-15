package main

import (
	"bufio"
	"fmt"
	"os"
)

//这本书真是给新手看的？搞不懂，我有基础看的都懵

func countLines(f *os.File, counts map[string]int) {
	//其实是和之前一样的用法，这里虽然是传值调用，但map作为一个指针数组（应该是？），指向的数组和传值指向的数组显然是一样的，里面的值变更也会影响外面
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
func main() {
	//这里同样是一个映射
	counts := make(map[string]int)
	files := os.Args[1:]
	//len是计算files长度，计算files长度又需要os.Args长度，目前还没明白切片到底是什么，先当个动态内存管理的数组看吧，就是可以根据大小自动扩容，初始大小为0，是C的NULL
	//所以如果文件名.go后面没有东西的话，files也会没有东西
	//如果这里的files为0，那么这就是之前算次数的逻辑
	if len(files) == 0 {
		//这里就是自己写的函数了，我习惯把函数放上面，圣经是放下面的
		//需要注意的是go中参数在后面，这里应该是一个传值调用，则os.Stdin是一个指针类型（*os.FIle)
		//这里的意思是，如果终端go后面没有东西，就等用户输入，下面转到函数内解释
		countLines(os.Stdin, counts)
	} else {
		//需要注意的是，这代表终端后方输入了其他东西
		//_.表示忽略这个值，因为go中用不到的变量会报错
		for _, arg := range files {
			//这里的Open就是c语言中的fopen，用于打开文件，arg表示文件的路径，如果不是路径就会报错（需要注意的是，路径也是字符串）
			//需要注意的是os.Open打开成功会返回文件地址和nil，错误会返回nil文件（就是NULL）和一个非nil的值
			f, err := os.Open(arg)
			if err != nil {
				//欸注意一下，Fprintf和Printf，printf是标准输入输出（也就是键盘），而Fprintf则更加灵活（文件夹，网络这种）
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
