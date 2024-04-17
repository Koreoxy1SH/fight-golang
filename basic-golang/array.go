package main

import "fmt"

func main() {

	var name [3]string
	name[0] = "Juki"
	name[1] = "budi"
	name[2] = "kuki"

	var number = [4]int{
		1,
		2,
		3,
		4,
	}

	var address = [2]string{
		"adrees1",
		"adrees2",
	}

	var kosong = []int{}
	var phone = [...]int{
		28213212,
		21312312,
		23123123,
	}

	fmt.Println(name)
	fmt.Println(name[1])
	fmt.Println(len(name))

	fmt.Println(number)
	fmt.Println(number[1])
	fmt.Println(address)
	fmt.Println(address[1])
	fmt.Println(phone)
}
