package main

import "fmt"

type Flags uint

const (
	FlagUp           Flags = 1 << iota // 1
	FlagBroadcast                      //10
	FlagLoopback                       //100
	FlagPointToPoint                   //1000
	FlagMulticast                      // 10000
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }                   //这里是判断末尾是否为1
func TurnDown(v *Flags)     { *v &^= FlagUp }                               //（这里是取地址）把v最后一位设为0
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }                         //把v第二位设为1
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 } //第二位或者第五位有一个为1
func main() {
	var v Flags = FlagMulticast | FlagUp //10001
	fmt.Printf("%b %t\n", v, IsUp(v))    // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
