package main

import "fmt"

func main() {

	traficLight := "merah222"

	switch traficLight {
	case "merah":
		fmt.Println("Harus berhenti")
	case "kuning":
		fmt.Println("Hati-hati")
	case "hijau":
		fmt.Println("Boleh Jalan")
	default:
		fmt.Println("Lampu rusak")
	}

	//Switch dengan short statement
	switch length := len(traficLight); length > 5 {
	case true:
		fmt.Println("Kata lebih dari lima")
	case false:
		fmt.Println("Kata tidak lebih dari lima")

	}

	//switch tanpa kondisi
	//kondisi di switch expression tidak wajib
	//jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut
	length := len(traficLight)
	switch {
	case length > 5:
		fmt.Println("kata lebih dari lima")
	case length < 5:
		fmt.Println("Kata tidak lebih dari lima")
	}
}
