package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	result := fmt.Sprintf("%s 下载完成", filename)
	fmt.Println(result)
	results <- result
}

func main() {
	fmt.Println("开始下载3个文件")
	var wg sync.WaitGroup
	results := make(chan string, 3)
	var s = []string{"file1.zip", "file2.pdf", "file3.mp4"}
	for _, b := range s {
		wg.Add(1)
		go download(b, &wg, results)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	wg.Wait()
	fmt.Println("所有文件下载完成")
}
