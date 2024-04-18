package main

import "fmt"

//Struct adalah tipe dapa seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
//namun jika kita ingin menambahka mehtod ke dalam Struct, sehingga seakan-akan sebuah struct memiliki function
//method adalah function

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is ", customer.Name)
}

func main() {
	budi := Customer{"Budi", "sada", 30}

	budi.sayHello("Udin")
}
