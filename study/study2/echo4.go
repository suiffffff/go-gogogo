package main

import (
	"flag"
	"fmt"
	"strings"
)

// 似乎是关于终端的一些东西
// 这里的flag可以当作标识名，将"n"当作了一个终端的标识，输入-n便可以激活
// 第二个值便是默认值，可更改
// 同时，这里的n会接受默认值。如果终端收到了-n的指令，就会将false改为true
// 最后一句似乎是固定写法，可以通过输入-h或者-help来显现
// 需要注意的是，这里接受的是指针
var n = flag.Bool("n", false, "omit trailing newline")

// 同理，这里用-s来激活
// ”“应该是可修改的，如果不使用-s，那么sep里面的值就应该是”“
// 最后一个同样是帮助信息
var sep = flag.String("s", "", "seperator")

func main() {
	//这里是解析终端输入的东西，不打这一行应该不会解析，flag开始遍历查找标识（-n，-s）
	//-n -s ”“ hello world
	flag.Parse()
	//string.Join显然拼接字符串
	//在flag.Parse收集完指令信息后，剩下的不认识的会被存储，Args就会返回剩下的字符串切片["hello" "world"]
	//这里会有一个版本特性，”“在go1.1后会被跳过，然后指向hello，所以此时src储存的就是hello，Join需要” “，那么便只会输出一个world
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Printf(*sep)
	//这里！表示相互，也就是终端未接受-n打印空格
	if !*n {
		fmt.Println()
	}
}
