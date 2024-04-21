package main

import "fmt"

//Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam methoda adalah pass by value
//sangat direkomendasikan menggunakan pointer di method, sehingg tidak boros memory karena harus selalu diduplikasi ketika memanggil method

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {

	budi := Man{"budi"}
	budi.Married()

	fmt.Println(budi.Name)

}
