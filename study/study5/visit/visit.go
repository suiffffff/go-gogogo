package visit

import "golang.org/x/net/html"

//这里出现了个很有趣的问题，那就是findlinks2无法直接使用findlinks1定义的函数
//为了解决这个问题，将visit函数封装成了一个包
//也可以将visit重命名
//至于为什么呢？因为go run findlinks2.go，只会运行一个文件，找不到visit定义

// *html.Node是个指向node结构体的指针，我们可以在html中看到node的结构，node显然是个树状结构
func Visit(links []string, n *html.Node) []string {
	//这里就和前端的html有关了，不多介绍，感兴趣可以自行了解
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//向下遍历递归
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}
