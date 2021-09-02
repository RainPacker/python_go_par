package main

import (
	"fmt"
	"go/constant"
)

// 自定义类型
type calc func(x, y int) int

func sum(x, y int) int {
	return x + y
}
func cal(x, y int, cb calc) int {
	return cb(x, y)
}
func show(a interface{}) {
	fmt.Println(a)
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
	show(1)
	show(constant.Bool)

	// 任意类型map
	var ma = make(map[string]interface{})
	ma["name"] = "zhang"
	ma["age"] = 1
	fmt.Println(ma)
	var anyStr interface{}
	anyStr = "11"
	// 断言 x.(type) 只能在swith中使用
	_, ok := anyStr.(int)
	if ok {
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")

	}

}
