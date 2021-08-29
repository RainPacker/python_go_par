package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func add(x int, y int) int {
	return x + y
}
func getUserInfo() (int, string) {
	return 29, "zhangyangyang"
}

func main() {
	fmt.Println(add(42, 13))
	var a, b int
	a = 10
	b = 11
	c := 12.131415926
	fmt.Println(a, b, c)
	fmt.Printf("c的类型是%T\n", c)
	fmt.Printf("c=%.2f c的二进制是：%b \n", c, c)
	fmt.Printf("a=%v b=%v c=%x ", a, b, c)
	var username string
	fmt.Println(username)
	var (
		d string
		e int
	)
	d = "zhang"
	d2 := "yang"
	d3 := fmt.Sprintf("%v-%v", d, d2)
	fmt.Println(d3)
	e = 1
	fmt.Println(d, e, unsafe.Sizeof(d), unsafe.Sizeof(e))
	f, g, h := 1, 2, "2"
	fmt.Printf("f的类型是%T g的类型是%T h 的类型是%T", f, g, h)
	//var age, name = getUserInfo()
	// 匿名变量 _
	var age, _ = getUserInfo()
	fmt.Println(age)
	// 常量
	const pi = 3.14
	fmt.Printf("pi=%v 类型是%T \n", pi, pi)
	// 计数 iota
	const y = iota
	fmt.Println(y)
	const (
		z, p = iota + 1, iota + 2
		_, _
		zz, pp
	)
	fmt.Println(z, p, zz, pp)
	// 切片
	var str = "a-b-c"
	println(str)
	split := strings.Split(str, "-")
	fmt.Println(split)
	fmt.Println(strings.Join(split, ","))
	var arr = []string{"a", "b", "c"}
	fmt.Println(arr)
	str2 := "-"
	fmt.Println(strings.Index(str, str2))
	fmt.Println(strings.LastIndex(str, str2))
	// 字符 byte
	var by = 'a'
	fmt.Println(unsafe.Sizeof(by))
	fmt.Printf("by=%v by=%c type is : %T \n", by, by, by)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c)\n", str[i], str[i])

	}
	var str3 = "你好golang"
	for _, v := range str3 {
		fmt.Printf("%v(%c) 类型：%T\n ", v, v, v)
	}
	byts := []byte(str)
	byts[0] = 'A'
	fmt.Println(string(byts))
	byts2 := []rune(str3) // 字符中有中文使用 rune
	byts2[0] = '真'
	fmt.Println(string(byts2))
	pistr := fmt.Sprintf("%f", pi)
	fmt.Println(pistr)
	fmt.Println(strconv.FormatInt(int64(e), 10))
	fmt.Println(strconv.FormatFloat(pi, 'f', 2, 64))
	fmt.Println(strconv.ParseInt("123", 10, 32))

}
