package main

import "fmt"

//Defer
//Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah function selesai di eksekusi
//Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
