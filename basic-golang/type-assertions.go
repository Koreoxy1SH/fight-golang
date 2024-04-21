package main

import "fmt"

//Type Assertions merupakan kemamapuan merubah tipe data menjadi tipe data yang diinginkan
//Fitur ini sering sekali digunakan ketika kita bertemu dengan tipe data interface kosong

func random() interface{} {
	return "OK"
}

func main() {
	var result any = random()
	//var changeDataType = result.(string)
	//fmt.Println(changeDataType)

	/**
	changeInt := result.(int)
	fmt.Println(changeInt)
	*/

	//Type Assertions Menggunakan switch
	//Saat salah Menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita\
	//agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unkonw type")
	}
}
