package main

import "fmt"

func main() {
	//Closure adalah kemampauan sebuah function berinterkasi dengan data-data disekitarnya dalam scope yang sama
	//harap menggunakan fitur closure ini dengan bijak saat kita membuat aplikasi
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
}
