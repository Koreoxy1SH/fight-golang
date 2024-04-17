package main

import "fmt"

// Function Type Declarations
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//Function tidak hanya bisa kita simpan di dalam  variable sebagai Value
	//namun juga bisa kita gunakan sebagai parameter untuk function lain

	sayHelloWithFilter("Budi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
