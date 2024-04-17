package main

import "fmt"

func getFullName() (string, string) {

	return "Budi", "Santoso"
}

func main() {
	//function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
	//untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function
	fmt.Println(getFullName())

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)

	//Menghiraukan Return value
	//Multiple return value wajib ditangkap semua valuenya
	//jika kita ingin menghiraukan return value tersebut, kita bissa menggunakan tanda _ (garis bawah)
	_, namaBelakang := getFullName()
	fmt.Println(namaBelakang)
}
