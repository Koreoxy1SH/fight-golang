//PACKAGE INITIALIZATION
//Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
//ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat functrion inisialisasi untuk membuka koneksi ke database
//untuk membuat function yg diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

//BLANK IDENTIFIER
//Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
//secara default, Golang akan komplan ketika ada package yg diimport namun tidak digunakan
//untuk menangani hal tersebut, ktia bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import

package main

import (
	"basic-golang/database"
	_ "basic-golang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
