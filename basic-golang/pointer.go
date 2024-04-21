package main

import "fmt"

//Pass by value
//Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
//artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirm adalah duplikasi value nya

//POINTER
//Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
//Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

//OPERATOR &
//untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variablenya

type Address struct {
	City, Province, Country string
}

type User struct {
	Name, Email, Password string
}

func main() {

	//Contoh code Pass by value
	address1 := Address{"Jakarata", "Jawa", "Indonesia"}
	address2 := address1

	address2.City = "Bali"
	fmt.Println(address2)
	fmt.Println(address1) //Tidak berubah value nya

	//Contoh code pointer
	user1 := User{"Budi", "budi@example.com", "12345"}
	user2 := &user1

	user2.Name = "Udin"
	fmt.Println(user2)
	fmt.Println(user1)
}
