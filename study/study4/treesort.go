package main

// 这里是一个很典型的二叉树，二叉树是什么就不介绍了，主要讲一下排列逻辑
type tree struct {
	value       int
	left, right *tree
}

// 我一向喜欢把需要的函数写到前面
func add(t *tree, value int) *tree {
	if t == nil {
		//new会创建一个结构体，把结构体里的值初始化后，把地址传给t，这里是在内存创建的，不存在出函数消失
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func appendValues(values []int, t *tree) []int {
	if t != nil {
		//这里会有递归，一致拿左边，显然左边是永远比上一个根节点小的
		values = appendValues(values, t.left)
		//接下来是拿中间，注意顺序是从最下开始网往上（递归）
		//需要注意的是，在左侧的一定比父节点小，在右侧的相当于在爷节点的左侧，也是比爷节点小的
		//整个逻辑链就是 父左- 父-父右（爷左）-爷-爷右然后递归左（也是父左）-父-父右
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func Sort(values []int) {

	//创建一个临时的变量来存储，这段代码本质上是个二叉树类型的排序方式
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}
func main() {
	//这里的逻辑有点多，顺便说一下
	//struct{}是一个空类型的结构体，本身可以当作一个类型
	//介绍一下set（集合），存在两个特性，1里面的元素是唯一的（不能重复）。2.主要用来判断“某个东西是否存在”
	//那么go中没有set，一般是用map模拟set，后面的值并不重要，主要是判断是否存在。
	//那么，bool肯定可以，这里的struct{}是属于比较邪门的做法。
	seen := make(map[string]struct{})
	s := "someString"
	//如果存在，跳过，如果不存在，改写标识
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
	}
}
