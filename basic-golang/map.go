package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "budi",
		"address": "jalan asdasd",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Naruto"
	book["author"] = "sukisa"
	book["wrong"] = "ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)

}
