package main

import "fmt"

//Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//Struct biasanya representasi data dalam program aplikasi yang kita buat
//data di struct disimpan dalam field
//sederhananya struct adalah kumpulan dari field

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var brian Customer
	brian.Name = "Brian"
	brian.Address = "Jalan kenangan"
	brian.Age = 30

	fmt.Println(brian)

	dudi := Customer{
		Name:    "dUDI",
		Address: "JLASJDSA",
		Age:     20,
	}

	fmt.Println(dudi)

	asep := Customer{"asep", "kenangan", 30}
	fmt.Println(asep)
}
