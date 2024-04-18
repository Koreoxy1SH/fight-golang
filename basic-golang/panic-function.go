package main

import "fmt"

//Panic Function adalah function yang bisa kita gunakan untuk menghentikan program
//Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
//saat panic function dipanggil, program akan terhenti, namun defer function tetap dieksekusikan

func endApp() {
	fmt.Println("App End")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(true)
}
