package main

import "fmt"

func main() {

	// Operasi perbandingan adalah operasi untuk membandingkan duah buah data
	// operasi perbandingan adalah operasi yg mengahasilkan nilai boolean (true or false)
	// jika hasil operasinya adalah true, maka nilainya adalah true
	// jika hasil operasinya adalah false, maka nilainya adalah false

	var name1 = "brian"
	var name2 = "brian"

	var result bool = name1 == name2
	fmt.Println(result)

	var address1 = "jl budi"
	var address2 = "jl merah"

	var merge bool = address1 != address2
	fmt.Println(merge)
}
