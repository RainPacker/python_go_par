package main

import "fmt"

type Animal struct {
	Name string
}

// Actons 定义行为接口
type Actons interface {
	run()
	setName(string) string
}

func (a Animal) run() {
	fmt.Printf("%v在运动\n", a.Name)
}

// 继承 animal
type Dog struct {
	Age int
	Animal
}

func (a Dog) run() {
	fmt.Printf("dog run ...-%v在运动\n", a.Animal.Name)
}
func (d *Dog) setName(s string) string {
	d.Animal.Name = s
	return d.Animal.Name
}

// (d Dog) 值接收者
func (d Dog) wang() {
	fmt.Printf("%v在汪汪\n", d.Name)
}

func main() {

	dog := &Dog{
		Age: 1,
		Animal: Animal{
			Name: "ww",
		},
	}
	fmt.Println(dog)
	dog.run()
	dog.wang()
	// 让 dog 实现 Actons 接口
	var action Actons = dog
	fmt.Println(action.setName("zyy1"))
	action.run()

}
