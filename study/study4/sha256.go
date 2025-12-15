package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//这里是哈希值，将x转化为一个32byte（也是256位）的数组（相当于数组存储了256二进制位，合起来组成这个哈希数）
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
