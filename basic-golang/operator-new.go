package main

import "fmt"

// Operator New
// Sebelumnya untuk membuat pointer dengan menggunakan operator &
// Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
// Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
type User struct {
	Name, Email, Password string
}

func main() {
	user1 := new(User)
	user2 := user1

	user2.Name = "Dudi"

	fmt.Println(user2)
	fmt.Println(user1)
}
