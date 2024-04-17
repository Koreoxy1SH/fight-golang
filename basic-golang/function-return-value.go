package main

import "fmt"

func getHello(name string) string {

	return "Hello " + name
}

func main() {
	//Function bisa mengembalikan data
	//untuk memberitahukan bahwa function mengembalikan data, kita harus menuliskan tip data
	//jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita haru mengembalikan data
	//untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya
	result := getHello("Brian")
	fmt.Println(result)
	fmt.Println(getHello("Santoso"))

}
