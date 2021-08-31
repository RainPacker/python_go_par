package main

import "fmt"

// 自定义类型
type calc func(x, y int) int

func sum(x, y int) int {
	return x + y
}
func cal(x, y int, cb calc) int {
	return cb(x, y)
}

func main() {
	var s calc
	s = sum
	fmt.Println(s(2, 3))
	fmt.Printf("s 的类型是%T\n", s)
	b := sum
	fmt.Printf("b的类型是%T\n", b)
	fmt.Println(cal(10, 19, s))
	var a = func() {
		print("我是匿名函数")
	}
	fmt.Println(a)
	a()
}
