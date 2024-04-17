package main

import "fmt"

func main() {
	//for loops perulangan

	counter := 1

	for counter <= 5 {
		fmt.Println("Perulangan ke ", counter)
		counter++
		//counter--
	}

	fmt.Println("Perulangan selesai")

	//For dengan statement
	//dalam For, kita bisa menambahkan statement, dimana terdapat 2 statement yg bisa tambahkan di for {
	//Init statement, yaitu statement sebelum for di eksekusi
	//Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan
	for decounter := 5; decounter >= 1; decounter-- {
		fmt.Println("Pengurangan ke ", decounter)
	}
	fmt.Println("Pengurangan selesai")

	//For Range
	//For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	//Data collection contohnya Array, Slice dan Map
	names := []string{"brian", "dudi", "kuki"}

	for i := 0; i < len(names); i++ { //ini mengakses secara manual
		fmt.Println(names[i])
	}

	for i, name := range names { //ini menggunakan for range
		fmt.Println("index = ", i, "Value = ", name)

	}

	for _, name := range names { //ini menggunakan for range
		fmt.Println(name)
	}
}
