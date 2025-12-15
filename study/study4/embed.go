package main

import "fmt"

type Point struct {
	X, Y int
}

//	type Circle struct {
//		Center Point
//		Radius int
//	}
//
//	type Wheel struct {
//		Circle Circle
//		Spokes int
//	}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	//#对于字符串能打印""，对于结构体能打印类型，对于空能打印nil或{}
	fmt.Printf("%#v\n", w)
}
