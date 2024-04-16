package main

import "fmt"

func main() {

	var name string

	name = "Dudi"
	fmt.Println(name)

	name = "Budi"
	fmt.Println(name)

	var address = "jl setia budi"
	var phone = 2132131
	fmt.Println(address)
	fmt.Println(phone)

	//:= adalah singkatan untuk membuat variable
	hobby := "Basketball"
	fmt.Println(hobby)

	hobby = "Soccer"
	fmt.Println(hobby)

	//Multiple variable
	var (
		orderList = "product list"
		status    = "done"
	)

	fmt.Println(orderList)
	fmt.Println(status)

}
