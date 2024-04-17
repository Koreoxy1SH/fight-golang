package main

import "fmt"

func sayGoodBye(name string) string {
	return "BYE-BYE " + name
}

func main() {

	result := sayGoodBye

	fmt.Println(result("Pujan"))

}
