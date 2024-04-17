package main

import "fmt"

func main() {
	//IF adalah salah satu kata kunci yg digunakan untuk percabangan
	//Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
	//hampir semua bahasa pemograman mendukung if expression

	word := "Helloxx"

	if word == "Hello" {
		fmt.Println(word, "Brian")
	} else if word == "Hola" {
		fmt.Println("hola")
	} else {
		fmt.Println("Word not found")
	}

	//IF dengan short statement
	//if mendukung short statement sebelum kondisi
	//hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi
	if length := len(word); length > 5 {
		fmt.Println("Kata terlalu panjang")
	} else {
		fmt.Println("Kata sudah benar")
	}
}

