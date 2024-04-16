package main

import "fmt"

func main() {

	//Constant adalah variable yg nilainya tidak bisa diubah lagi setelah pertama kali diberikan nilai pada variable tersebut
	//saat pembuatan constant, wajib langsung menginisialisasikan datanya.

	const firstName string = "budi"
	const lasstName = "santoso"

	const (
		address = "Jl dudadsa"
		phone   = 12321312
	)

	fmt.Println(firstName, lasstName)
	fmt.Println(address, phone)

	//Error tidak bisa diubah lagi
	//const firstName = "juju"

}
