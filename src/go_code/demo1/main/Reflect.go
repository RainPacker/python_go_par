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
	field := el.Elem().Field(0)
	fmt.Println(field)
	name := field.Name
	ty := field.Type
	tag := field.Tag.Get("json")
	fmt.Println(name, ty, tag)

	filed1, ok := el.Elem().FieldByName("Age")
	if ok {
		fmt.Println(filed1.Name, filed1.Type, filed1.Tag)
	}
	method, ok := el.Elem().MethodByName("PrintInfo")
	fmt.Println(ok)
	if ok {
		fmt.Println(method)
	}
	var val = reflect.ValueOf(i)
	val.MethodByName("PrintInfo").Call(nil)
	//if el.Elem().Kind() == reflect.Struct || el.Kind() == reflect.Struct {
	//	field := el.Field(0)
	//	fmt.Println(field)
	//}
	// 通过反射调整 方法 ChangeName
	var params []reflect.Value
	params = append(params, reflect.ValueOf("zhangyangyang"))
	// 传入带参数的切片
	val.MethodByName("ChangeName").Call(params)
	fmt.Println("方法数量：", el.NumMethod())
}

type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	addr string `json:"addr,omitempty"`
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
	fmt.Printf("修改后的 名称=%v", s.Name)
}

func (s Student) PrintInfo() {
	fmt.Printf("name=%v age =%v addr=%v", s.Name, s.Age, s.addr)
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
	GetFiled(&stu)
}
