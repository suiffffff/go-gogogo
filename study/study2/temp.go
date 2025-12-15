package main

//如何引用一个包？首先需要 go mod init study2（命名，这是模块名） 然后就可以通过模块名-包名导入
import (
	"fmt"
	"os"
	"strconv"
	"study2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprint(os.Stderr, "cf:%v", err)
			os.Exit(1)
		}
		//如果要引用的话就需要如下的写法，很类似结构体的访问
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s=%s %s=%s", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
