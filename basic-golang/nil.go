package main

import "fmt"

//Biasanya di dalam bahasa pemograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
//Berbeda dengan GoLang, di GoLang saa kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
//Namun di GoLang ada data nil, nilainya yaitu data kosong.
//Nil Sendiri hanya bisa digunakan di bebrapa tipe data, seperti interface, function, map, slice, pointer dan channel.

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("tahu")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}

	test := NewMap("")
	fmt.Println(test["name"])

}
