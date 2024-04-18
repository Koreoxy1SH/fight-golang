package main

import "fmt"

//Recover adalah function yang bisa kita gunakan untuk menangkap data panic
//dengan Recover proses panic akan terhenti, sehingga program akan tetap berjalan

func endApp() {
	fmt.Println("App End")

	//CARA YANG BENAR MENGGUNAKAN RECOVER
	message := recover()
	fmt.Println("Terjadi error ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	//CARA YANG SALAH MENGGUNAKAN RECOVER
	//message := recover()
	//fmt.Println("terjadi errro ", message)
}

func main() {
	runApp(true)
}
