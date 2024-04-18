package main

import "fmt"

//Go-Lang bukanlah bahasa pemograman yang berorientasi objek.
//Biasanya dalam pemograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemograman tersebut,
//contoh nya di java ada java.lang.object
//untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
//Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
//Interface kosong, juga memiliki type alias bernama any

func Ups() interface{} {
	return 1
}

func Kosong() any {
	return 1
}

func main() {

	kosong := Ups()
	fmt.Println(kosong)

	emptyy := Kosong()
	fmt.Println(emptyy)
}
