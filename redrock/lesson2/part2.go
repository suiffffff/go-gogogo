package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	results <- fmt.Sprintf("%s 下载完成", filename)
}

func main() {
	fmt.Println("开始下载3个文件")
	var wg sync.WaitGroup
	results := make(chan string, 3)
	var s = []string{"file1.zip", "file2.pdf", "file3.mp4"}
	for _, file := range s {
		wg.Add(1)
		go download(file, &wg, results)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	wg.Wait()
	for res := range results {
		fmt.Println(res)
	}
	fmt.Println("所有文件下载完成")
}
