//Di bahasa pemograman lain, biasanya ada kata kunci yg bisa digunakan untuk menentukan access modifier terhadap suatu function atu variable
//Di golang, untuk menentukan access modifier, cukup dengan nama function atau variable
//jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain

package main

import (
	"basic-golang/helper"
	"fmt"
)

func main() {

	fmt.Println(helper.App)
	name := helper.SayHello("budi")
	fmt.Println(name)

	//ERROR TIDAK DITEMUKAN
	//bye := helper.sayBye("sadsa")

}
