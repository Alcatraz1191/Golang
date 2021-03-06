package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//In a simple factory pattern, we make a function which will return an instance of Person struct.
func NewPerson(name string, age int) Person {
	p := Person{
		Name: name,
		Age:  age,
	}
	return p
}

func main() {
	p := Person{Name:"Shrey"}
	fmt.Println(p)
	//Regular declaration and definition would be something like
	//a := Person{Name : "Shrey", Age : 24}. Don't have to worry about forgetting to initialize any value.
	a := NewPerson("Shrey", 24)
	fmt.Println(a)
}
