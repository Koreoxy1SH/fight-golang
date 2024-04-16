package main

import "fmt"

func main() {
	// Type declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada.
	// Type declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

	type NumberNIM int
	type (
		matkul string
		total  NumberNIM
		status bool
	)

	var NIM NumberNIM = 213123213
	var name_matkul matkul = "PEMOGRAMAN"
	var active status = true
	var total_matkul total = 2

	fmt.Println(NIM, name_matkul, active, total_matkul)
	fmt.Println(NumberNIM(123213))

}
