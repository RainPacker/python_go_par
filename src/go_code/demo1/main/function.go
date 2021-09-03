package main

import (
	"fmt"
	"go/constant"
	"time"
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

// 匿名返回
func fun1() int {
	var a int = 0
	defer func() {
		a++
	}()
	return a
}
func fun2() (a int) {
	//var a int = 0
	defer func() {
		a++
	}()
	return a
}
func fun3() {
	defer func() {
		// 可以捕获到异常信息 recover
		err := recover()
		if err != nil {
			fmt.Println("error=", err)
		}
	}()
	panic("fun3抛出一个异常。。。")
}

type Usb interface {
	start()
	stop()
}
type Phone struct {
	Name string
}
type Computer struct {
}

func (computer Computer) work(usb Usb) {
	usb.start()
	usb.stop()
}

func (phone Phone) stop() {
	fmt.Println("手机闭关...")
}

func (phone Phone) start() {
	fmt.Println("小米手机启动...")
}

type Camer struct {
}

func (c Camer) start() {
	fmt.Println("相机启动...")
}

func (c Camer) stop() {
	fmt.Println("相机关闭...")
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
	// 延迟执行
	fmt.Println(a)
	a()
	fmt.Println("")
	fmt.Println(fun1())
	fmt.Println(fun2())
	fun3()
	fmt.Println(time.Now())
	now := time.Now()
	formatStr := now.Format("2006-01-02 15:04:05")
	fmt.Println(formatStr)
	fmt.Println(now.Unix())
	times := time.Unix(1630395929, 0)
	fmt.Println(times)
	fmt.Printf("%d-%02d-%d %v:%v:%v", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 日期str 转 uinix
	str := "2020-01 12:20:20"
	var temp string = "2006-01-02 15:04:05"
	timeInput, _ := time.ParseInLocation(temp, str, time.Local)
	fmt.Println("time", timeInput.Unix())
	// 定时器
	ticker := time.NewTicker(time.Second)
	fmt.Println(ticker)
	// ticker.C
	n := 10
	for t := range ticker.C {
		fmt.Println(n)
		n--

		if n < 0 {
			ticker.Stop()
			break
		}
		fmt.Println(t)
	}
	time.Sleep(time.Second) //休眠
	n1 := &n
	fmt.Println(n1, *n1)
	var phone = Phone{
		Name: "小米",
	}
	var usb Usb
	usb = phone // 手机继承了 Usb 接口
	usb.start()
	//usb.stop()
	var computer = Computer{}
	var carmer = Camer{}
	computer.work(phone)
	computer.work(carmer)
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
