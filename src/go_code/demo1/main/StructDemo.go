package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v在运动\n", a.Name)
}

// 继承 animal
type Dog struct {
	Age int
	Animal
}

func (d Dog) wang() {
	fmt.Printf("%v在汪汪\n", d.Name)
}

func main() {
	dog := Dog{
		Age: 1,
		Animal: Animal{
			Name: "ww",
		},
	}
	fmt.Println(dog)
	dog.run()
	dog.wang()
}
