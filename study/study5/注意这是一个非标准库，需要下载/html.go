package 注意这是一个非标准库_需要下载

// go get golang.org/x/net/html
// 这是这个库的部分内容
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32

type Attribute struct {
	Key, Val string
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)
