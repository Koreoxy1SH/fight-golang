package main

import "fmt"

//Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
//sebuah interface berisikan definisi-definisi method
//biasanya interface digunakan sebagai kontrak

type HashName interface {
	getName() string
}

func SayHello(value HashName) {
	fmt.Println("Hello ", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {

	person := Person{Name: "JUki"}
	SayHello(person)

	animal := Animal{Name: "Harimau"}
	SayHello(animal)
}
