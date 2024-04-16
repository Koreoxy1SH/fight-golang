package main

import "fmt"

func main() {

	//Convert data number
	var tipe32 int32 = 32768
	var tipe64 int64 = int64(tipe32)
	var tipe16 int16 = int16(tipe32)

	fmt.Println("tipe data 32 = ", tipe32)
	fmt.Println("tipe data 64 = ", tipe64)
	fmt.Println("tipe data 16 = ", tipe16)

	//convert data number to string
	var name = "budi santoso"
	var index1 = name[0]
	var index2 = name[1]
	var convertString1 = string(index1)
	var convertString2 = string(index2)

	fmt.Println(convertString1, convertString2)
}
