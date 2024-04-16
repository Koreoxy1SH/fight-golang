package main

import "fmt"

func main() {
	//Operasi boolean
	// && = Dan
	// || = atau
	// ! = kebalikan

	var nilaiAkhir = 90
	var absensi = 70

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusNilaiAbsensi bool = absensi >= 80

	var result = lulusNilaiAkhir && lulusNilaiAbsensi

	fmt.Println(result)
}
