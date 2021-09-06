package main

import (
	"fmt"
	"reflect"
)

// 使用反射修改值
func chanVal(i interface{}) {
	var el = reflect.ValueOf(i)
	fmt.Println("el", el.Elem())
	fmt.Println(el.Elem().Kind())
	if el.Elem().Kind() == reflect.Int64 {
		el.Elem().SetInt(999)
	}
	if el.Elem().Kind() == reflect.String {
		el.Elem().SetString("我是修改后值")
	}
	if el.Elem().Kind() == reflect.Struct || el.Kind() == reflect.Struct {
		field := el.Field(0)
		fmt.Println(field)
	}
}
func GetFiled(i interface{}) {
	var el = reflect.TypeOf(i)
	fmt.Println(el)
	field := el.Field(0)
	fmt.Println(field)
	//if el.Elem().Kind() == reflect.Struct || el.Kind() == reflect.Struct {
	//	field := el.Field(0)
	//	fmt.Println(field)
	//}
}

type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	addr string `json:"addr,omitempty"`
}

func main() {
	var a int64 = 100
	of := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		fmt.Println("int", v.Int())
	}
	fmt.Println(v)
	fmt.Println(v.Int())
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(of.Name()) // 类别名称
	fmt.Println(of.Kind()) // 种类
	chanVal(&a)
	fmt.Printf("a=%v", a)
	var b = "你好 go lang"
	chanVal(&b)
	fmt.Println(b)

	var po *string = &b
	fmt.Println(po)
	*po = "zxxxx"
	fmt.Println(b)
	var stu = Student{
		Name: "zyy",
		Age:  30,
		addr: "wx",
	}
	GetFiled(stu)
}
