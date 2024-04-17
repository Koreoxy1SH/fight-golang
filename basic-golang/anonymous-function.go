package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {

	blackList := func(name string) bool {
		return name == "anjing"
	}

	registerUser("brian", blackList)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
