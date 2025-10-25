package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//make可以制造切片与映射（这里是映射），现在我们可以把map当做一个数组，[]里面则替换掉了常规意义上的数组下标 如arr[0]指第一个元素，而现在map[string]指向第一个元素，同样的，int替换掉了数组里面表示的元素（也可以当作元素类型？）
	//如果很难理解的话，可以想象一下数学的函数，y=kx，每一个x([])，有一个对应的y(int)这也就是为什么叫映射的原因
	counts := make(map[string]int)
	//bufio.Newcanner也是一个新的函数，这里很难用c语言解释，我们可以想象input作为参数进入了bufio.Newscanner，这个函数里面提供了一些功能函数，input可以再次作为参数进入
	//os.Stdin是一个显而易见的参数，这个表示标准输入输出，推测和stdio.h应该差不多，是一个固定格式
	input := bufio.NewScanner(os.Stdin)
	//Scan()显然也是一个函数，会读取下一行输入，这里循环需要ctrl+d结束
	for input.Scan() {
		//Text()同样也是一个函数，会获取Scan()读取的内容，然后在往外看，就是counts[],变成了一个和c语言类似的数组引用，之后的++自然是对数组内元素的++，而不是string
		counts[input.Text()]++

	}
	//这里的range遍历的是映射，[]的对应第一个变量，也就是line对应string，n同理对应int
	for line, n := range counts {
		if n > 1 {
			//所以这个函数实现的功能就是获得输入数的次数，以及输入了什么
			fmt.Printf("%d\t%s\n", n, line)
		}

	}
}
