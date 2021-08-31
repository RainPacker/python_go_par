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

// 无返回值
func demo(x ...int) []int {
	fmt.Println(x)
	return x
}
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}

	return slice
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
	fmt.Println(e << 2)
	if e > 2 {
		fmt.Println("e  是大于2的")
	} else {
		fmt.Println("不大于2")
	}
	if score := 80; score > 95 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良")
	} else if score >= 70 {
		fmt.Println("不错")
	} else {
		fmt.Println("要努力了")
	}
	// golang 中没有 while
	i := 1
	for i <= 10 {
		fmt.Println("golang")
		i++
	}
	// 打印乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v \t", i, j, i*j)
		}

		fmt.Println("")

	}
	// forrang 的使用
	title := "你好golang"
	for k, v := range title {
		fmt.Println(k, v)
	}
	switch num := 7; num {

	case 1, 3, 5, 7, 9:
		fmt.Println("是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("是偶数")

	}

	// 切片
	var arr2 = []string{"a", "b", "c"}
	// 数组
	var arr3 = [...]string{"a", "b", "c"}
	fmt.Println(arr2, arr3)
	arr4 := arr3[:]  // 全部
	arr4 = arr3[1:2] // 全部
	fmt.Println(arr4)
	fmt.Printf("长度是%d 容量是:%d", len(arr4), cap(arr4))
	fmt.Println(append(arr2, "d"))
	fmt.Println(append(arr2, "d"))
	// map
	var userIfo = make(map[string]string)
	userIfo["age"] = "20"
	userIfo["name"] = "zyy"
	fmt.Println(userIfo)
	var usrInfo2 = map[string]string{
		"name": "zyy",
		"age":  "30",
		"h":    "168",
	}
	usrInfo2["x"] = "xxxx"
	usrInfo2["x"] = "222"
	fmt.Println(usrInfo2["name"])
	for k, v := range usrInfo2 {
		fmt.Printf("%v=%v \n", k, v)
	}

	delete(usrInfo2, "h")
	// 查找
	s, ok := usrInfo2["h"]
	fmt.Println(s, ok)
	var info = make(map[string][]string)

	info["hobby"] = []string{
		"java",
		"golang",
	}
	fmt.Println(info)
	demo(1, 2, 3)
	var slice = []int{1, 3, 6, 5, 2, 7, 9, 4}

	fmt.Println(sortIntAsc(slice))

}
