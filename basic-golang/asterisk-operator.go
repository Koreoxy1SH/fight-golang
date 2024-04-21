package main

import "fmt"

//Operator *
//Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
//semua variable yang mengacu ke data yang sama tidak akan berubah
//jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

type User struct {
	Name, Email, Password string
}

func main() {

	//Contoh code pointer
	user1 := User{"Budi", "budi@example.com", "12345"}
	user2 := &user1

	user2.Name = "Udin"
	fmt.Println(user2)
	fmt.Println(user1)

	*user2 = User{"Kuji", "kuji@gmail.com", "21312"}
	fmt.Println(user2)
	fmt.Println(user1)
}
