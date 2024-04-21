package main

import "fmt"

//Pointer di function
//Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirm ke function tersebut.
//Oleh karean itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
//HAl ini membuat variabel menjadi aman, karena tidak akan bisa berubah
//namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut.
//untuk melakukan ini, kita juga bisa menggunakan pointer di fucntion/
//untuk menjadikan sebuah paramter sebagai pointer, kita bisa menggunakan operator * di parameternya

type User struct {
	Name, Email, Password string
}

func ChangeNameToUdin(user *User) {
	user.Name = "Udin"
}

func main() {

	user := User{}
	ChangeNameToUdin(&user)

	fmt.Println(user)
}
