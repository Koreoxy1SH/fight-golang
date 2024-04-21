// PACKAGE
// Package adalah tempat yg bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
// Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
// Packge sendiri sebenarnya hanya direktori folder di sistem operasi kita
package main

//IMPORT
//Secara standar, file Go-Lang hanya bisa mengakses file go-lang lainnya yang berbeda dalam package yg sama
//jika kita ingin mengakses file go-lang yang berbeda diluar package, maka kita bisa menggunakan import
import (
	"basic-golang/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("budi")
	fmt.Println(result)
}
