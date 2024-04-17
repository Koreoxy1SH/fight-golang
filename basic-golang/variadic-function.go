package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	total := sumAll(20, 20, 20)
	fmt.Println(total)
	fmt.Println("=================")
	fmt.Println(sumAll(10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
