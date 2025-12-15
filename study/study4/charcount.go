package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	//utf8.UTFMax是一个常量，表示UTF8最大可能占的字节数
	//不知道说没说，go里面的字符串基本都是UTF8编码
	//+1就是5,这里是一个数组类型 []int，+1是很常用的对应关系。不用0开头
	//需要注意的是，这里面的，1代表1字节的utf8，2是2字节的，以此类推
	var utflen [utf8.UTFMax + 1]int
	//这里是统计的无效字符个数
	invalid := 0
	//NewReader自然是带缓冲的读取了
	//那么和NewScanner有什么不同呢？
	//scanner是按行读取的，最大缓存是64kib，Text来获得文本内容
	//reader相当于把你输入的所有东西存起来，然后按需读取，Read.Byte(读取一个字节）ReadRune(读取一个rune)ReadString('\n')（读取一个换行）
	//并且没有缓存限制
	//
	in := bufio.NewReader(os.Stdin)
	for {
		//r是rune，读取的字符，n表示int，占的字节数，err是错误
		r, n, err := in.ReadRune()
		//EOF，End of files，文件结尾，不是文件结尾就继续读
		if err == io.EOF {
			break
		}
		//这里是遇到错误了
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		//这里是统计坏字节，当readrune无法解析的时候，会返回一个特殊的替换字符unicode。replacementchar（），n是1（一个字节）
		//invalid+1，然后跳过本次循环
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		//执行到这里就是没问题，没问题就是正常字符
		counts[r]++
		//这里就是在统计这个是几字节的utf8码
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	//以下为使用map
	fmt.Printf("\nchar\tcount (高到低)\n")
	type charCount struct {
		r rune
		n int
	}
	var sortedCount []charCount
	for r, n := range counts {
		sortedCount = append(sortedCount, charCount{r, n})
	}
	sort.Slice(sortedCount, func(i, j int) bool {
		return sortedCount[i].n > sortedCount[j].n
	})
	for i, sc := range sortedCount {
		if i >= 10 {
			break
		}
		fmt.Printf("%q\t%d\n", sc.r, sc.n)
	}
}

//那么还可以说说的是，这里的暂存需要回车发送，而回车是带一个“\n”的
//所有输出一定会有一个一字节的换行
//而ctrl+d只是用来结束输出，并不会发送
