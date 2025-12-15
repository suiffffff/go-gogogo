package main

import (
	"fmt"
	"os"
	"strconv"
	//导入包是需要模块化的，模块化是什么？其实就是go.mod这一个文件里module后面的名字
	//模块化后是包的名字，需要注意的是不是文件路径，这里的tempconv文件名改成什么都可以，我们引用的是package后面的名字
	"study2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		//str显然是个字符串函数，conv表示转换，float表示类型，64表示精度，这个函数的意思是把字符串转化为浮点数
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			//这里的警告是因为fprintf是需要返回值的，如果需要改可以改成以下形式
			//_,_=fmt.Fprintf(os.Stderr, "cf:%v", err)
			//这里也可以扩展一下Fprintf与Printf的区别
			//F代表文件，Printf显然是从标准输出流输出（os.Stdout），和Fprintf相比，F代表着输出方向
			//Fprintf可以向例如（os.Stderr）（标准错误流）这样的地方输入
			//因此不难发现，Printf其实等价于Fprintf（os.Stdout ...)
			fmt.Fprintf(os.Stderr, "cf:%v", err)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s=%s,%s=%s", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
