package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type Person struct {
	Name string `json:"name"`
	age  int
	addr string
}

// PrintPersonInfo 给 Person 添加行为方法
func (p Person) PrintPersonInfo() {
	fmt.Printf("name=%v age =%v addr=%v\n", p.Name, p.age, p.addr)
}
func (p *Person) SetAge(age int) {
	p.age = age
}

func main() {
	var a = new(int) // 创建一个 int 类型的指针
	fmt.Printf("a 的值是%v 类型是%T %v", a, a, unsafe.Sizeof(a))
	var p1 Person
	p1.Name = "zyy"
	p1.age = 29
	p1.addr = "addr"

	fmt.Println("")
	fmt.Println(p1)
	fmt.Printf("p1的值是%v 类型是%T \n", p1, p1)
	fmt.Printf("p1的值是%#v 类型是%T", p1, p1)

	var p2 = new(Person)
	p2.Name = "rain"
	p2.age = 11
	p2.addr = "wx"
	fmt.Println("###########")
	fmt.Println("p2=", p2)
	var p3 = &Person{}
	p3.Name = "zyy"
	fmt.Println(p3)
	p2.PrintPersonInfo()
	var p4 = Person{
		Name: "zhang",
		age:  11,
		addr: "11",
	}
	p4.SetAge(30)
	p4.PrintPersonInfo()
	// 将p4 转换成 json 字符 注 只用是大写的（Public）才能序列化）
	jsonByte, _ := json.Marshal(p4)
	jsons := string(jsonByte)
	fmt.Println(jsons)
	var p5 Person
	jsonStr := `{"name":"zhangyang-p5"}`
	// json字符 转成 struct
	err := json.Unmarshal([]byte(jsonStr), &p5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p5)

}
