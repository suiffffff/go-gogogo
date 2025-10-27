package main

import (
	"fmt"
)

func main() {
	var a, b int
	var c string
	for {
		fmt.Scanf("%d %s %d", &a, &c, &b)
		//不好用啊这个，感觉不如newscanner
		if c == "exit" {
			break
		}
		switch c {
		case "+":
			fmt.Printf("%d + %d = %d\n", a, b, a+b)
		case "-":
			fmt.Printf("%d - %d = %d\n", a, b, a-b)
		case "*":
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
		case "/":
			if b == 0 {
				fmt.Println("除数不能为零")
			} else {
				fmt.Printf("%d / %d = %d\n", a, b, a/b)
			}
		default:
			fmt.Println("非标准格式。请重新输入")
		}
	}
}
