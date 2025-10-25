package main

import "fmt"

// 突然开始简单了起来，这是一个摄氏度和华氏度的转换
// 所以这才是新手的开头吗
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
